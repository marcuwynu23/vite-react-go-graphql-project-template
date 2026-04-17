import { useMutation } from "@tanstack/react-query";
import type { FormEvent } from "react";
import { useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { api } from "@/lib/axios";
import { useAuthStore } from "@/store/useStore";

type LoginResponse = { token: string };
type LoginMutationResponse = {
  data: {
    login: LoginResponse;
  };
};

export function LoginPage() {
  const navigate = useNavigate();
  const location = useLocation();
  const setToken = useAuthStore((s) => s.setToken);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const from =
    (location.state as { from?: { pathname: string } })?.from?.pathname ||
    "/dashboard";

  const login = useMutation({
    mutationFn: async () => {
      const { data } = await api.post<LoginMutationResponse>("/graphql", {
        query: `
          mutation Login($input: LoginInput!) {
            login(input: $input) {
              token
            }
          }
        `,
        variables: {
          input: {
            email,
            password,
          },
        },
      });
      return data.data.login;
    },
    onSuccess: (data) => {
      setToken(data.token);
      navigate(from, { replace: true });
    },
  });

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    login.mutate();
  }

  function skipDemo() {
    setToken("demo-token");
    navigate(from, { replace: true });
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-neutral-50 px-4 dark:bg-neutral-950">
      <div className="w-full max-w-sm rounded-xl border border-neutral-200 bg-white p-6 shadow-sm dark:border-neutral-800 dark:bg-neutral-900">
        <h1 className="text-lg font-semibold tracking-tight">Sign in</h1>
        <p className="mt-1 text-sm text-neutral-600 dark:text-neutral-400">
          Use the API login or continue with a local demo session.
        </p>
        <form className="mt-6 space-y-4" onSubmit={onSubmit}>
          <div>
            <label
              className="block text-xs font-medium text-neutral-700 dark:text-neutral-300"
              htmlFor="email"
            >
              Email
            </label>
            <input
              id="email"
              type="email"
              autoComplete="username"
              className="mt-1 w-full rounded-md border border-neutral-300 bg-white px-3 py-2 text-sm outline-none ring-neutral-400 focus:ring-2 dark:border-neutral-600 dark:bg-neutral-950"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="you@example.com"
            />
          </div>
          <div>
            <label
              className="block text-xs font-medium text-neutral-700 dark:text-neutral-300"
              htmlFor="password"
            >
              Password
            </label>
            <input
              id="password"
              type="password"
              autoComplete="current-password"
              className="mt-1 w-full rounded-md border border-neutral-300 bg-white px-3 py-2 text-sm outline-none ring-neutral-400 focus:ring-2 dark:border-neutral-600 dark:bg-neutral-950"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
            />
          </div>
          {login.isError && (
            <p className="text-sm text-red-600 dark:text-red-400">
              Sign in failed. Try demo sign-in or check the API.
            </p>
          )}
          <Button type="submit" className="w-full" disabled={login.isPending}>
            {login.isPending ? "Signing in…" : "Sign in"}
          </Button>
        </form>
        <Button
          type="button"
          variant="outline"
          className="mt-3 w-full"
          onClick={skipDemo}
        >
          Demo sign-in (no API)
        </Button>
      </div>
    </div>
  );
}
