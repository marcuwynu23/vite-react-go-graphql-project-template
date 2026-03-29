import { QueryClient } from "@tanstack/react-query";

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 60 * 1000,
      retry: (failureCount, error) => {
        if (error && typeof error === "object" && "response" in error) {
          const status = (error as { response?: { status?: number } }).response
            ?.status;
          if (status && status >= 400 && status < 500 && status !== 408) {
            return false;
          }
        }
        return failureCount < 2;
      },
      refetchOnWindowFocus: false,
    },
    mutations: {
      retry: false,
    },
  },
});
