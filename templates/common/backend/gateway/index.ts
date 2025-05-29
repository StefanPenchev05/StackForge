import express from "express";
import { createProxyMiddleware } from "http-proxy-middleware";
import routes from "./routes.json";
import authMiddleware from "./middlewares/auth";

// Types
import { RouteDefinition } from "./types/RouteDefinition";

const app = express();
const PORT = process.env.PORT || 3000;

Object.entries(routes as Record<string, RouteDefinition>).forEach(([path, config]) => {
  const middlewares = [];

  if (config.authRequired) {
    middlewares.push(authMiddleware);
  }

  middlewares.push(
    createProxyMiddleware({
      target: config.target,
      changeOrigin: true,
      pathRewrite: (pathReq) => pathReq.replace(path, ""),
    })
  );

  app.use(path, ...middlewares);
});

app.listen(PORT, () => {
  console.log(`Gateway listening on port ${PORT}`);
});
