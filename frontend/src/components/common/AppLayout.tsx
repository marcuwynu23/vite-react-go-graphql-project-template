import type { ReactNode } from "react";
import { Link } from "react-router-dom";
import { useAuthStore } from "@/store/useStore";

export function AppLayout({ children }: { children: ReactNode }) {
  const setToken = useAuthStore((s) => s.setToken);

  return (
    <div className="min-h-screen bg-neutral-50 text-neutral-900 dark:bg-neutral-950 dark:text-neutral-50">
      <header className="border-b border-neutral-200 bg-white dark:border-neutral-800 dark:bg-neutral-900">
        <div className="mx-auto flex max-w-5xl items-center justify-between px-4 py-3">
          <Link to="/dashboard" className="font-semibold tracking-tight">
            Dashboard
          </Link>
          <button
            type="button"
            className="rounded-md border border-neutral-300 px-3 py-1.5 text-sm hover:bg-neutral-100 dark:border-neutral-600 dark:hover:bg-neutral-800"
            onClick={() => setToken(null)}
          >
            Sign out
          </button>
        </div>
      </header>
      <main className="mx-auto max-w-5xl px-4 py-8">{children}</main>
    </div>
  );
}
