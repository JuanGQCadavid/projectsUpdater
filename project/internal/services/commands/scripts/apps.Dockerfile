FROM openjdk:8-jre
#
ARG SERVICE_NAME='factaggregator'
ARG APP_NAME='FactAggregator'
#
ARG APP_HOME_DIR=/var/factiva/Apps/${APP_NAME}
ARG DIST_DIR='/DIST'
ARG TMP_DIR=$APP_HOME_DIR/tmp
# common
ARG NEWRELIC_VERSION=7.5.0
ARG RUN_USER=metadata
ARG RUN_GROUP=metadata
ARG GID=1000
 
 
# '/var/factiva/tmp'
ENV APP_ENV='stag'
ENV GRANDEOURSE_ENV='grandeourse_stag'
ENV DGRANDEOURSE_DATACENTER='virginia'
ENV APP_PROPERTY_FILE='factaggregator.properties'
ENV APP_MAIN_CLASS='FactAggregatorApp'
 
RUN mkdir -p $DIST_DIR \
    $TMP_DIR \
    $APP_HOME_DIR/newrelic/logs
 
WORKDIR $APP_HOME_DIR/logs
WORKDIR $APP_HOME_DIR/live   
 
WORKDIR $APP_HOME_DIR
COPY ./target/platformenrichment-${SERVICE_NAME}-*-spring-boot.jar app.jar
COPY ./target/shared-resources/properties_new properties
COPY ./target/shared-resources/data data
COPY ./src/main/resources/bin/start.sh start.sh
 
RUN groupadd -r ${RUN_GROUP} &&\
    useradd -g ${RUN_GROUP} -d ${APP_HOME_DIR} -s /bin/bash ${RUN_USER}
 
RUN chown -R ${RUN_USER}:${RUN_GROUP} $APP_HOME_DIR
## validate user permission
RUN ls -lah $APP_HOME_DIR
RUN su -c 'touch ${APP_HOME_DIR}/this.will.fail' ${RUN_USER}
 
RUN chmod -R 777 ${APP_HOME_DIR} && chmod -R 777 ${TMP_DIR}
 
USER $RUN_USER