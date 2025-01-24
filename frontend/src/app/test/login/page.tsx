"use client";
import React, { useState } from "react";
import { useAuth } from "@/app/hooks/auth/useAuth"; // useAuth をインポート

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { handleLogin, isLoading, error } = useAuth(); // useAuth フックを使用
  const [message, setMessage] = useState("");

  const onSubmit = async () => {
    try {
      await handleLogin(email, password); // handleLogin を呼び出し
      setMessage("ログイン成功しました！");
    } catch (err) {
      console.error("ログイン失敗:", err);
      setMessage("ログインに失敗しました。");
    }
  };

  return (
    <div style={{ maxWidth: "400px", margin: "auto", padding: "20px" }}>
      <h1>ログイン</h1>
      <input
        type="email"
        placeholder="メールアドレス"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        style={{ width: "100%", padding: "10px", marginBottom: "10px" }}
      />
      <input
        type="password"
        placeholder="パスワード"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        style={{ width: "100%", padding: "10px", marginBottom: "10px" }}
      />
      <button
        onClick={onSubmit}
        disabled={isLoading}
        style={{
          width: "100%",
          padding: "10px",
          backgroundColor: isLoading ? "#ccc" : "#4CAF50",
          color: "white",
          border: "none",
          cursor: isLoading ? "not-allowed" : "pointer",
        }}
      >
        {isLoading ? "ログイン中..." : "ログイン"}
      </button>
      {error && (
        <p style={{ marginTop: "20px", color: "red" }}>エラー: {"エラー"}</p>
      )}
      {message && (
        <p
          style={{
            marginTop: "20px",
            color: message.includes("成功") ? "green" : "red",
          }}
        >
          {message}
        </p>
      )}
    </div>
  );
};

export default LoginPage;
