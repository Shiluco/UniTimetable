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
    ? localStorage.getItem("authToken") // トークンを取得
    : null;

  const headers = {
    ...options.headers,
    ...(token ? { Authorization: `Bearer ${token}` } : {}), // トークンがある場合のみ追加
  };

  console.log("Request URL:", url);
  console.log("Request Options:", { ...options, headers });

  const response = await fetch(url, { ...options, headers });

  if (response.status === 401) {
    window.location.href = "/login";
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
