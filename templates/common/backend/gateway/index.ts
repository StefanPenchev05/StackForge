import express from "express";
import path from "path";
import routes from "./routes.json";
import authMiddleware from "./middlewares/auth";
import { createProxyMiddleware } from "http-proxy-middleware";

// Types
import { RouteDefinition } from "./types/RouteDefinition";

const app = express();
const PORT = process.env.PORT || 3000;

Object.entries(routes as Record<string, RouteDefinition>).forEach(([routePath, config]) => {
  const middlewares = [];

  if (config.authRequired) {
    middlewares.push(authMiddleware);
  }

  if(config.customMiddleware) {
    for(const mw of config.customMiddleware) {
        try {
            const mwPath = path.resolve(__dirname, mw.path);
            const mwFunc = require(mwPath).default;
            middlewares.push(mwFunc);
            console.log(`✅ Loaded middleware "${mw.target}" for ${routePath}`);
        } catch(err) {
            console.warn(`⚠️ Failed to load middleware "${mw.target}" at path ${mw.path}:`, (err as Error).message);
        }
    }
  }

  middlewares.push(
    createProxyMiddleware({
      target: config.target,
      changeOrigin: true,
      pathRewrite: (pathReq) => pathReq.replace(routePath, ""),
    })
  );

  app.use(routePath, ...middlewares);
});

app.listen(PORT, () => {
  console.log(`Gateway listening on port ${PORT}`);
});
