package com.platform.userservice.services;

import com.platform.userservice.config.Redission;
import org.redisson.api.RBucket;
import org.redisson.api.RScript;
import org.redisson.api.RedissonClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class RedisServices {
    @Autowired
    private RedissonClient redissonClient;

    // ✅ Set a value with optional TTL
    public <T> void setValue(String key, T value, long ttlSeconds) {
        RBucket<T> bucket = redissonClient.getBucket(key);
        bucket.set(value, ttlSeconds, TimeUnit.SECONDS);
    }

    // ✅ Get a value
    public <T> T getValue(String key) {
        RBucket<T> bucket = redissonClient.getBucket(key);
        return bucket.get();
    }

    // ✅ Delete a key
    public boolean deleteKey(String key) {
        return redissonClient.getBucket(key).delete();
    }

    // ✅ Run a Lua script
    public Object evalLuaScript(String script, List<String> keys) {
        return redissonClient.getScript().eval(
                RScript.Mode.READ_WRITE,
                script,
                RScript.ReturnType.VALUE,
                keys
        );
    }
}
