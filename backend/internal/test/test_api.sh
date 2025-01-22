#!/bin/bash

# ベースURL
BASE_URL="http://localhost:8080/api"

# ユーザー登録
echo "Testing user registration..."
REGISTER_RESPONSE=$(curl -s -X POST "${BASE_URL}/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }')
echo $REGISTER_RESPONSE | jq '.'

# ログイン
echo -e "\nTesting login..."
LOGIN_RESPONSE=$(curl -s -X POST "${BASE_URL}/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }')
echo $LOGIN_RESPONSE | jq '.'

# トークンを抽出
TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.token')

# ユーザー一覧の取得
echo -e "\nTesting get users..."
curl -s "${BASE_URL}/users" \
  -H "Authorization: Bearer ${TOKEN}" | jq '.'

# メールアドレスで検索
echo -e "\nTesting email search..."
curl -s "${BASE_URL}/users/search/email?q=test" \
  -H "Authorization: Bearer ${TOKEN}" | jq '.'

# 名前で検索
echo -e "\nTesting name search..."
curl -s "${BASE_URL}/users/search/name?q=Test" \
  -H "Authorization: Bearer ${TOKEN}" | jq '.'