-- データベースが存在しない場合のみ作成
SELECT 'CREATE DATABASE unitimetable WITH OWNER postgres'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'unitimetable')\gexec
