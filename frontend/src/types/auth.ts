import { User } from "./user";

export type LoginResponse = {
  status: string; // レスポンスステータス (例: success, error)
  message: string; // ユーザーへの通知メッセージ
  error_detail?: string; // オプショナル: エラーの詳細情報（エラー時のみ）
  data: {
    accessToken: string; // アクセストークン
    refreshToken: string; // リフレッシュトークン
    token_type: string; // 例: Bearer
    expires_in: number; // トークンの有効期限（秒）
    user: User; // ユーザー情報
  };
};
