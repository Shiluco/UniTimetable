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
    "day_of_week": 1,
    "time_slot": 1,
    "subject": "プログラミング演習",
    "location": "情報処理演習室"
  }')
echo $SCHEDULE_RESPONSE | jq '.'

# 作成した時間割のIDを取得
SCHEDULE_ID=$(echo $SCHEDULE_RESPONSE | jq -r '.id')

# 投稿作成のテスト
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
POST_ID=$(echo $POST_RESPONSE | jq -r '.id')

# 投稿への返信作成のテスト
echo -e "\n${GREEN}Testing reply creation...${NC}"
REPLY_RESPONSE=$(curl -s -X POST "${BASE_URL}/posts" \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "Content-Type: application/json" \
  -d "{
    \"parent_post_id\": ${POST_ID},
    \"content\": \"返信テストです\"
  }")
echo $REPLY_RESPONSE | jq '.'

# 投稿の取得テスト
echo -e "\n${GREEN}Testing get post...${NC}"
curl -s "${BASE_URL}/posts/${POST_ID}" \
  -H "Authorization: Bearer ${TOKEN}" | jq '.'

# 返信の取得テスト
echo -e "\n${GREEN}Testing get replies...${NC}"
curl -s "${BASE_URL}/posts/${POST_ID}/replies" \
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