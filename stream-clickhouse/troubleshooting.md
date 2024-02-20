# 404 Not Found

## There is no handle /<Table Name>

Make sure Clickhouse config.xml file <http_server> setting allows interaction through HTTP. For example, valid config should looks like below:
```
<clickhouse>
    <!-- Other configuration settings -->

    <http_server>
        <listen_host>::</listen_host>
        <listen_port>8123</listen_port>
        <timezone>UTC</timezone>

        <!-- Allow access to all endpoints -->
        <endpoint default_handler="true" compression_disable="true">
            <base>/</base>
        </endpoint>
    </http_server>
</clickhouse>
```

Note: If you don't provide config.xml file, Clickhouse generates it's own preconfigured config!