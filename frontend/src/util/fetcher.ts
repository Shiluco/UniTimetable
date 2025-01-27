export const Fetcher = async <T>(
  endpoint: string, // エンドポイント
  options: RequestInit & { authRequired?: boolean } = { method: "GET", authRequired: true }
): Promise<T> => {
  const baseURL = "http://localhost:8080/api/"; // 環境変数からベースURLを取得
  if (!baseURL) {
    throw new Error("Base URL is not defined. Check your environment variables.");
  }

  const url = `${baseURL}${endpoint}`; // ベースURLとエンドポイントを結合

  const token = options.authRequired
    ? localStorage.getItem("accessToken") // トークンを取得
    : null;
  console.log("Fetcher - Token:", token);

  const headers = {
    ...options.headers,
    ...(token ? { Authorization: `Bearer ${token}` } : {}), // トークンがある場合のみ追加
  };

  console.log("Fetcher - Request URL:", url);
  console.log("Fetcher - Request Method:", options.method);
  console.log("Fetcher - Request Headers:", headers);
  if (options.body) {
    console.log("Fetcher - Request Body:", options.body);
  }

  const response = await fetch(url, { ...options, headers });

  // レスポンスのステータスをログ出したい場合
  console.log("Fetcher - Response Status:", response.status);

  if (response.status === 401) {
    throw new Error("Unauthorized: Please log in.");
  } else if (response.status === 403) {
    window.location.href = "/login";
    throw new Error("Forbidden: You do not have access to this resource.");
  } else if (response.status === 404) {
    window.location.href = "/login";
    throw new Error("Not Found: The requested resource could not be found.");
  }

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.message || "API request failed");
  }

  return response.json() as Promise<T>;
};
