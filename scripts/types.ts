export enum CacheType {
  /**
   * Cache type for Redis
   */
  REDIS = 'redis',
}

export enum StoreType {
  /**
   * Store type for Redis
   */
  REDIS = 'redis',
}

export interface CacheConfig {
  /**
   * Cache type
   */
  type: CacheType.REDIS;

  /**
   * Redis store configuration
   */
  store: StoreConfig;
}

export interface StoreConfig {
  /**
   * Redis host
   */
  host: string;

  /**
   * Redis port
   */
  port: number;

  /**
   * Redis password
   */
  password?: string;

  /**
   * Redis database
   */
  db?: number;
}