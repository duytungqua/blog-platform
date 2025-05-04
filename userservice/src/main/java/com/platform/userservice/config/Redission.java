package com.platform.userservice.config;

import org.redisson.api.RedissonClient;
import org.redisson.config.Config;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;
import java.io.InputStream;
import java.io.ObjectInputFilter;

@Configuration
public class Redission {
    @Bean(destroyMethod = "shutdown")
    public RedissonClient   redissonClient() throws IOException {
        // Load config from classpath
        InputStream configStream = getClass().getClassLoader().getResourceAsStream("redisson-config.yaml");
        ObjectInputFilter.Config config = Config.fromYAML(configStream);
        return Redission.create(config);
    }
}
