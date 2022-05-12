FROM tomcat:8.5.78-jdk8
 
##### envs
ENV APP_SERVICE_NAME='keyphrase'
ENV AXIS2_HOME=$CATALINA_HOME/webapps/axis2
ENV AXIS2_DEP_DIR=${AXIS2_HOME}/WEB-INF
ENV AXIS2_SERVICE_HOME=${AXIS2_DEP_DIR}/${APP_SERVICE_NAME}
ENV AXIS2_SERVICE_STATUS_HOME=${CATALINA_HOME}/webapps/keyphrasestatus
ENV RUN_USER=tomcat
ENV RUN_GROUP=tomcat
ENV GID=1000
 
ENV CATALINA_OPTS="-Xms2g -Xmx2g"
ENV JAVA_OPTS="-Djava.util.logging.config.file=$CATALINA_HOME/conf/logging.properties"
ENV JAVA_OPTS="$JAVA_OPTS -DGRANDEOURSE_DATACENTER=virginia"
ENV JAVA_OPTS="$JAVA_OPTS -Djava.util.logging.manager=org.apache.juli.ClassLoaderLogManager"
ENV JAVA_OPTS="$JAVA_OPTS -Djava.io.tmpdir=$CATALINA_HOME/temp -DGRANDEOURSE_ENV=int"
# ENV JAVA_OPTS="$JAVA_OPTS -Djavax.net.ssl.trustStore=$JAVA_HOME/jre/lib/security/cacerts"
# ENV JAVA_OPTS="$JAVA_OPTS -Djavax.net.ssl.trustStorePassword=changeit"
ENV JAVA_OPTS="$JAVA_OPTS -DCATALINA.BASE=$CATALINA_HOME"
ENV JAVA_OPTS="$JAVA_OPTS -DCATALINA.home=$CATALINA_HOME"
 
WORKDIR ${CATALINA_HOME}
COPY configs/axis2.zip ${CATALINA_HOME}/webapps/
RUN  unzip -oj ${CATALINA_HOME}/webapps/axis2.zip -d ${CATALINA_HOME}/webapps/ && \
     unzip ${CATALINA_HOME}/webapps/axis2.war -d ${CATALINA_HOME}/webapps/axis2 && \
     rm -rf ${CATALINA_HOME}/webapps/axis2.zip && \
     rm -rf ${CATALINA_HOME}/webapps/axis2.war &&\
     rm -rf ${CATALINA_HOME}/webapps/*.txt
 
## copy tomcat libs from dependencies
WORKDIR ${CATALINA_HOME}/lib
COPY target/workdir/apache-tomcat/lib/* .
COPY target/extras/lib/guava-25.0-jre.jar guava-25.0-jre.jar
 
## copy keyphrase soap lib from dependencies
WORKDIR ${AXIS2_SERVICE_HOME}/lib
COPY target/workdir/apache-tomcat/webapps/axis2/WEB-INF/keyphrase/lib/* .
 
## copy mainjar lib from dependencies
WORKDIR ${AXIS2_DEP_DIR}/services
COPY target/extras/mainjar-1.0.jar mainjar-1.0.jar
# <include>com.netvention.apps:mainjar</include>
# keep a backup for applications that could need a pristine from configure-webservice.rb
RUN rm -f mainjar-1.0.jar.bak && mv mainjar-1.0.jar mainjar-1.0.jar.bak
 
## copy keyphrasestatus lib from dependencies
WORKDIR ${AXIS2_SERVICE_STATUS_HOME}/WEB-INF/lib
COPY target/workdir/apache-tomcat/webapps/keyphrasestatus/WEB-INF/* .
# <include>org.json:*</include>
COPY target/extras/StatusServlet.jar StatusServlet.jar
 
## copy keyphrase properties
WORKDIR ${AXIS2_SERVICE_HOME}/properties
COPY ./configs/keyphrase_ws.properties keyphrase_ws.properties
COPY ./target/extras/properties/tagger.properties tagger.properties
COPY ./target/extras/properties/aggregator.properties aggregator.properties
COPY ./target/extras/properties/keyphrase.properties keyphrase.properties
COPY ./target/extras/properties/keyphrase_ws_monitor.properties keyphrase_ws_monitor.properties
COPY ./configs/kp_log4j.properties kp_log4j.properties
 
## copy tomcat files
WORKDIR ${CATALINA_HOME}/webapps
COPY ./target/extras/apache-tomcat/webapps/keyphrasestatus keyphrasestatus
COPY ./target/extras/apache-tomcat/webapps/KeyPhrase KeyPhrase
 
## copy keyphrase soap data files
WORKDIR ${AXIS2_SERVICE_HOME}/data
COPY ./target/extras/data/xml/DistDoc xml/DistDoc
COPY ./target/extras/data/xml/keyphrase-input.xsd xml/keyphrase-input.xsd
COPY ./target/extras/data/xml/keyphrase-output.xsd xml/keyphrase-output.xsd
COPY ./target/extras/data/tagger tagger
COPY ./target/extras/data/aggregator aggregator
COPY ./target/extras/data/keyphrase keyphrase
COPY ./target/extras/data/webservice/keyphrase webservice/keyphrase
COPY ./target/extras/data/unicodeTable.txt .
 
#Just creating the "logs" directory
WORKDIR ${CATALINA_HOME}/webapps/axis2/WEB-INF/keyphrase/logs
 
#### configure-webservice.rb
# create a link DIR of the $APP_SERVICE_NAME in $CATALINA_HOME
RUN ln -s $AXIS2_DEP_DIR/$APP_SERVICE_NAME $CATALINA_HOME/$APP_SERVICE_NAME
# place the service.xml file next to mainjar-1.jar
RUN mv $AXIS2_DEP_DIR/$APP_SERVICE_NAME/data/webservice/$APP_SERVICE_NAME/META-INF $AXIS2_DEP_DIR/services/
# place the lib file next to $APP_SERVICE_NAME.jar
RUN mv $AXIS2_DEP_DIR/$APP_SERVICE_NAME/lib $AXIS2_DEP_DIR/services/
# rename $APP_SERVICE_NAME.jar to $APP_SERVICE_NAME.aar
# RUN rm -f ${AXIS2_DEP_DIR}/mainjar-1.jar.bak && cp ${AXIS2_HOME}/services/mainjar-* ${AXIS2_HOME}/mainjar-1.jar.bak
RUN cp $AXIS2_DEP_DIR/services/mainjar-1.0.jar.bak $AXIS2_DEP_DIR/services/$APP_SERVICE_NAME.aar
# update the mainjar-1.jar with services.xml
WORKDIR $AXIS2_DEP_DIR/services
RUN /usr/local/openjdk-8/bin/jar uf $APP_SERVICE_NAME.aar META-INF/services.xml
# update the mainjar-1.jar with library files
RUN /usr/local/openjdk-8/bin/jar uf $APP_SERVICE_NAME.aar lib/*
# delete lib/* files
RUN rm -rf lib
# delete services.xml
RUN rm -rf META-INF
 
ENV APP_SERVICE_NAME='keyphrase'
ENV AXIS2_HOME=$CATALINA_HOME/webapps/axis2
ENV AXIS2_DEP_DIR=${AXIS2_HOME}/WEB-INF
ENV AXIS2_SERVICE_HOME=${AXIS2_DEP_DIR}/${APP_SERVICE_NAME}
 
#### enable ssl on tomcat
# config needs a ssl directory to keep security keys
WORKDIR ${CATALINA_HOME}/ssl
# platform-enrichment-keyphrase.fdev.dowjones.net-b64.cer == localhost-rsa-cert-b64.cer
# platform-enrichment-keyphrase.fdev.dowjones.net.pem == localhost-rsa-cert.pem
COPY ./configs/ssl/platform-enrichment-keyphrase.fdev.dowjones.net-b64.cer server.cer
COPY ./configs/ssl/platform-enrichment-keyphrase.fdev.dowjones.net.pem server.pem
COPY ./configs/ssl/localhost-rsa-key.pem localhost-rsa-key.pem
COPY ./configs/ssl/localhost-rsa-key.pem localhost-rsa-key.pem
COPY ./configs/ssl/rds-ca-2019-root.pem rds-ca-2019-root.pem
RUN echo yes | $JAVA_HOME/bin/keytool -importcert -alias MySQLCACert -file rds-ca-2019-root.pem -keystore $JAVA_HOME/jre/lib/security/cacerts -storepass changeit
 
RUN groupadd -r ${RUN_GROUP} &&\
    useradd -g ${RUN_GROUP} -d ${CATALINA_HOME} -s /bin/bash ${RUN_USER}
 
RUN chown -R ${RUN_USER}:${RUN_GROUP} $CATALINA_HOME
## validate user permission
RUN ls -lah $CATALINA_HOME
RUN su -c 'touch $CATALINA_HOME/work/this.will.fail' ${RUN_USER}
 
USER ${RUN_USER}
WORKDIR ${CATALINA_HOME}
EXPOSE 3180 8443 8080