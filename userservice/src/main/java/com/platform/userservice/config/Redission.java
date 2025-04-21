package com.platform.userservice.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class Redission {
    @Bean(destroyMethod = "shutdown")
    public RedissonClient redissonClient() throws IOException {
        // Load config from classpath
        InputStream configStream = getClass().getClassLoader().getResourceAsStream("redisson-config.yaml");
        Config config = Config.fromYAML(configStream);
        return Redisson.create(config);
    }
}
