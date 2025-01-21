export const Fetcher = async <T>(
  endpoint: string, // エンドポイント
  options: RequestInit & { authRequired?: boolean } = { authRequired: true }
): Promise<T> => {
  const baseURL = process.env.NEXT_PUBLIC_API_URL; // 環境変数からベースURLを取得

  if (!baseURL) {
    throw new Error(
      "Base URL is not defined. Check your environment variables."
    );
  }

  const url = `${baseURL}${endpoint}`; // ベースURLとエンドポイントを結合

  const token = options.authRequired
    ? localStorage.getItem("authToken") // トークンを取得
    : null;

  const headers = {
    ...options.headers,
    ...(token ? { Authorization: `Bearer ${token}` } : {}), // トークンがある場合のみ追加
  };

  const response = await fetch(url, { ...options, headers });

  if (response.status === 401) {
    throw new Error("Unauthorized: Please log in.");
  } else if (response.status === 403) {
    throw new Error("Forbidden: You do not have access to this resource.");
  } else if (response.status === 404) {
    throw new Error("Not Found: The requested resource could not be found.");
  }

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.message || "API request failed");
  }

  return response.json() as Promise<T>;
};
