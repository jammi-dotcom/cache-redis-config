interface CacheOptions {
  expire: number;
  maxlifetime: number;
  saveHandler?: string;
}

interface RedisOptions {
  host: string;
  port: number;
  password?: string;
  db?: number;
}

interface Config {
  cache?: CacheOptions;
  redis?: RedisOptions;
}

declare const config: Config;
export default config;