import { useQuery } from "@tanstack/react-query";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { api } from "@/lib/axios";

type HealthResponse = { status: string; service: string };
type HealthQueryResponse = {
  data: {
    health: HealthResponse;
  };
};

export function DashboardPage() {
  const health = useQuery({
    queryKey: ["health"],
    queryFn: async () => {
      const { data } = await api.post<HealthQueryResponse>("/graphql", {
        query: `
          query Health {
            health {
              status
              service
            }
          }
        `,
      });
      return data.data.health;
    },
  });

  return (
    <div className="space-y-4">
      <div>
        <h1 className="text-2xl font-semibold tracking-tight">Dashboard</h1>
        <p className="mt-1 text-sm text-neutral-600 dark:text-neutral-400">
          Frontend and backend are wired through{" "}
          <code className="rounded bg-neutral-100 px-1 py-0.5 text-xs dark:bg-neutral-800">
            VITE_API_URL
          </code>{" "}
          and the dev proxy for{" "}
          <code className="rounded bg-neutral-100 px-1 py-0.5 text-xs dark:bg-neutral-800">
            /graphql
          </code>
          .
        </p>
      </div>
      <Card>
        <CardHeader>
          <CardTitle>Backend health</CardTitle>
          <CardDescription>GraphQL health query via Axios + React Query</CardDescription>
        </CardHeader>
        <CardContent>
          {health.isLoading && (
            <p className="text-sm text-muted-foreground">Checking…</p>
          )}
          {health.isError && (
            <p className="text-sm text-destructive">
              Could not reach the API. Start the Go server on port 8080.
            </p>
          )}
          {health.data && (
            <pre className="overflow-x-auto rounded-lg bg-muted p-3 text-xs">
              {JSON.stringify(health.data, null, 2)}
            </pre>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
