package org.ex.distributed_computing.config;

import org.springframework.cache.Cache;
import org.springframework.cache.CacheManager;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.cache.RedisCacheManager;
import org.springframework.data.redis.connection.RedisConnectionFactory;

@Configuration
public class CacheConfig {

  @Bean
  public CacheManager cacheManager(RedisConnectionFactory redisConnectionFactory) {
    return RedisCacheManager.create(redisConnectionFactory);
  }

  @Bean
  public Cache writerCache(CacheManager cacheManager) {
    return cacheManager.getCache("writerCache");
  }
}
