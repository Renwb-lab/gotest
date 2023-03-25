
docker run --name elasticsearch -p 9200:9200  -p 9300:9300 \
-e "discovery.type=single-node" \
-e ES_JAVA_OPTS="-Xms84m -Xmx512m" \
-e log4j.configurationFile="" \
-v /Users/wenbin.ren/DockerFile/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml \
-v /Users/wenbin.ren/DockerFile/elasticsearch/data:/usr/share/elasticsearch/data \
-v /Users/wenbin.ren/DockerFile/elasticsearch/plugins:/usr/share/elasticsearch/plugins \
-d elasticsearch:7.12.0

docker run -d -p 5601:5601 -v /Users/wenbin.ren/DockerFile/kibana/config/kibana.yml:/usr/share/kibana/config/kibana.yml --restart=always --name kibana kibana:7.12.0

curl -X GET "http://192.168.1.239:9200/_cat/health?v=true&pretty"

