export type LoginResponse = {
  status: string;
  message: string;
  error_detail: string;
  data: {
    accessToken: string;
    refreshToken: string;
    token_type: string; // 例: Bearer
    expires_in: number; // トークンの有効期限（秒）
    error_detail: string;
    user: {
      user_id: number;
      username: string;
      email: string;
      department_id: number;
      major_id: number;
      grade: number;
    };
  };
};
