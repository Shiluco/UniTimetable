import { Fetcher } from "@/util/fetcher";
import { LoginResponse } from "@/types/auth";
import { RegisterResponse } from "@/types/auth";

export const login = async (email: string, password: string): Promise<LoginResponse> => {
  const response = await Fetcher<LoginResponse>("auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
    headers: { "Content-Type": "application/json" },
    authRequired: false,
  });
  return response;
};

export const register = async (name: string, email: string, password: string): Promise<RegisterResponse> => {
  const response = await Fetcher<RegisterResponse>("auth/register", {
    method: "POST",
    body: JSON.stringify({ name, email, password }),
    headers: { "Content-Type": "application/json" },
    authRequired: false,
  });
  return response;
};
