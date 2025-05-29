export type RouteDefinition = {
    target: string;
    authRequired?: boolean;
    customMiddleware?: CustomMiddleware[]
}

type CustomMiddleware = {
    path: string;
    target: string
}