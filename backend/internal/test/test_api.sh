#!/bin/bash

# ベースURL
BASE_URL="http://localhost:8080/api"

# 色の定義
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

# テスト結果の追跡
TESTS_PASSED=0
TESTS_FAILED=0

# テスト関数
run_test() {
    local test_name=$1
    local command=$2
    local expected_status=$3

    echo -e "\n${GREEN}Running test: ${test_name}${NC}"
    
    # コマンドを実行し、結果とステータスコードを取得
    local response=$(eval $command)
    local status=$?
    
    # レスポンスを表示
    echo "Response:"
    echo $response | jq '.'
    
    # テスト結果の判定
    if [ $status -eq $expected_status ]; then
        echo -e "${GREEN}✓ Test passed${NC}"
        TESTS_PASSED=$((TESTS_PASSED + 1))
    else
        echo -e "${RED}✗ Test failed${NC}"
        TESTS_FAILED=$((TESTS_FAILED + 1))
    fi
}

# ユーザー登録のテスト
echo -e "\n${GREEN}Testing user registration...${NC}"
REGISTER_RESPONSE=$(curl -s -X POST "${BASE_URL}/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "テストユーザー",
    "email": "test@example.com",
    "password": "password123"
  }')
echo $REGISTER_RESPONSE | jq '.'

# トークンを取得
TOKEN=$(echo $REGISTER_RESPONSE | jq -r '.token')

# ログインのテスト
echo -e "\n${GREEN}Testing login...${NC}"
LOGIN_RESPONSE=$(curl -s -X POST "${BASE_URL}/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }')
echo $LOGIN_RESPONSE | jq '.'

# トークンを更新
TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.token')

# 時間割作成のテスト
echo -e "\n${GREEN}Testing schedule creation...${NC}"
SCHEDULE_RESPONSE=$(curl -s -X POST "${BASE_URL}/schedules" \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{
    "day_of_week": 0,
    "time_slot": 1,
    "subject": "プログラミング演習",
    "location": "情報処理演習室"
  }')
echo $SCHEDULE_RESPONSE | jq '.'

# 作成した時間割のIDを取得
SCHEDULE_ID=$(echo $SCHEDULE_RESPONSE | jq -r '.schedule_id')

# 通常の投稿作成のテスト
echo -e "\n${GREEN}Testing post creation...${NC}"
POST_RESPONSE=$(curl -s -X POST "${BASE_URL}/posts" \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "Content-Type: application/json" \
  -d "{
    \"content\": \"テスト投稿です\",
    \"schedule_id\": ${SCHEDULE_ID}
  }")
echo $POST_RESPONSE | jq '.'

# 作成した投稿のIDを取得
POST_ID=$(echo $POST_RESPONSE | jq -r '.post_id')

# 返信作成のテスト
echo -e "\n${GREEN}Testing reply creation...${NC}"
REPLY_RESPONSE=$(curl -s -X POST "${BASE_URL}/posts" \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "Content-Type: application/json" \
  -d "{
    \"parent_post_id\": ${POST_ID},
    \"content\": \"返信テストです\"
  }")
echo $REPLY_RESPONSE | jq '.'

# 単一の投稿取得テスト
echo -e "\n${GREEN}Testing get single post...${NC}"
curl -s "${BASE_URL}/posts/${POST_ID}" \
  -H "Authorization: Bearer ${TOKEN}" | jq '.'

# JSONレスポンスを処理する関数
process_response() {
    local response=$1
    if echo "$response" | jq -e . >/dev/null 2>&1; then
        if [[ $(echo "$response" | jq -r 'type') == "object" ]]; then
            if [[ $(echo "$response" | jq 'has("error")') == "true" ]]; then
                echo "$response" | jq '.'
            else
                echo "$response" | jq '.posts // .'
            fi
        else
            echo "$response" | jq '.'
        fi
    else
        echo "Invalid JSON response: $response"
    fi
}

# 返信一覧の取得テスト
echo -e "\n${GREEN}Testing get replies...${NC}"
GET_REPLIES_RESPONSE=$(curl -s "${BASE_URL}/posts?parent_id=${POST_ID}" \
  -H "Authorization: Bearer ${TOKEN}")
process_response "$GET_REPLIES_RESPONSE"

# 時間割に関連する投稿一覧の取得テスト
if [ ! -z "$SCHEDULE_ID" ]; then
    echo -e "\n${GREEN}Testing get posts by schedule...${NC}"
    GET_SCHEDULE_POSTS_RESPONSE=$(curl -s "${BASE_URL}/posts?schedule_id=${SCHEDULE_ID}" \
      -H "Authorization: Bearer ${TOKEN}")
    process_response "$GET_SCHEDULE_POSTS_RESPONSE"
fi

# 投稿の更新テスト
echo -e "\n${GREEN}Testing post update...${NC}"
curl -s -X PUT "${BASE_URL}/posts/${POST_ID}" \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "更新されたテスト投稿です"
  }' | jq '.'

# 投稿の削除テスト
echo -e "\n${GREEN}Testing post deletion...${NC}"
curl -s -X DELETE "${BASE_URL}/posts/${POST_ID}" \
  -H "Authorization: Bearer ${TOKEN}" | jq '.'

# テスト結果の表示
echo -e "\n${GREEN}Test Summary:${NC}"
echo "Passed: ${TESTS_PASSED}"
echo "Failed: ${TESTS_FAILED}"

# 終了コードの設定
if [ $TESTS_FAILED -eq 0 ]; then
    exit 0
else
    exit 1
fi