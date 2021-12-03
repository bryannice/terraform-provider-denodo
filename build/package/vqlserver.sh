#!/bin/bash
# -----------------------------------------------------------------------------
#  Environment variable JAVA_HOME must be set and exported
# -----------------------------------------------------------------------------

if [ "$VDB_CONFIGURATION_FILE" == "" ]
then
    VDB_CONFIGURATION_FILE=VDBConfiguration.properties
fi

DENODO_JRE_HOME="/opt/denodo/jre"
DENODO_JAVA_HOME="/opt/denodo"
DENODO_JRE11_OPTIONS="-Djava.locale.providers=COMPAT,SPI -Dcom.datastax.driver.FORCE_NIO=true"

if [ "/opt/denodo" != "" ]
then
    if [ -e "/opt/denodo/jre/bin/java" ]
    then
        JAVA_BIN="/opt/denodo/jre/bin/java"
    else
        JAVA_BIN="/opt/denodo/bin/java"
    fi
fi
if [ ! -e "$JAVA_BIN" ]
then
    if [ -d "$DENODO_JRE_HOME" ]
    then
        JAVA_BIN=$DENODO_JRE_HOME/bin/java
    fi
fi
if [ ! -e "$JAVA_BIN" ]
then
    if [ -e "$JAVA_HOME/jre/bin/java" ]
    then
        JAVA_BIN="$JAVA_HOME/jre/bin/java"
    else
        JAVA_BIN="$JAVA_HOME/bin/java"
    fi
fi

DENODO_LAUNCHER_CLASSPATH="/opt/denodo/lib/denodo-commons-launcher-util.jar"
if [ "$DENODO_EXTERNAL_CLASSPATH" != "" ]
then
    DENODO_LAUNCHER_CLASSPATH=$DENODO_LAUNCHER_CLASSPATH:$DENODO_EXTERNAL_CLASSPATH
fi

export LD_LIBRARY_PATH="/opt/denodo/bin:/opt/denodo/dll/db:/opt/denodo/extensions/thirdparty/native"

if [ -e "$JAVA_BIN" ]
then
    case "$1" in
      startup)
         if [ "$2" == "-tray" ]
         then
           echo "Starting Denodo Virtual DataPort Server with tray icon ... (press ENTER in case shell prompt was not automatically returned)"
           "$JAVA_BIN" -DDENODO_APP="Denodo VDP Server 8.0" -Xmx4096m -XX:+DisableExplicitGC -XX:+UseG1GC -XX:ReservedCodeCacheSize=256m  $DENODO_JRE11_OPTIONS $DENODO4E_DEBUG_CONF -Djava.system.class.loader=com.denodo.util.launcher.DenodoClassLoader \
             -DconfFile=$VDB_CONFIGURATION_FILE \
             -Djava.util.logging.manager=org.apache.logging.log4j.jul.LogManager \
             -Ddenodo.rmi.server.hostname.default \
             -Djavax.xml.xpath.XPathFactory:http://java.sun.com/jaxp/xpath/dom=com.sun.org.apache.xpath.internal.jaxp.XPathFactoryImpl \
             -Djavax.xml.transform.TransformerFactory=com.sun.org.apache.xalan.internal.xsltc.trax.TransformerFactoryImpl \
             -Djavax.xml.parsers.SAXParserFactory=com.sun.org.apache.xerces.internal.jaxp.SAXParserFactoryImpl \
             -Djavax.xml.parsers.DocumentBuilderFactory=com.sun.org.apache.xerces.internal.jaxp.DocumentBuilderFactoryImpl \
             -Djavax.xml.stream.XMLOutputFactory=com.ctc.wstx.stax.WstxOutputFactory \
             -D"com.microsoft.tfs.jni.native.base-directory=/opt/denodo/dll/vdp/tfs" \
             -DverboseMode=false \
             -Djdk.serialFilter="!org.mozilla.**;!org.apache.commons.**;" \
             -Ddenodo.launcher.excluded.jars="windows" \
             -classpath "$DENODO_LAUNCHER_CLASSPATH" \
             com.denodo.util.launcher.Launcher com.denodo.vdb.vdbinterface.server.VDBManagerImpl \
             --conf "/opt/denodo/conf" \
             --conf "/opt/denodo/resources/apache-tomcat/conf" \
             --conf "/opt/denodo/conf/vdp" \
             --lib "/opt/denodo/lib/vdp-client-core" \
             --lib "/opt/denodo/lib/vdp-server-core" \
             --lib "/opt/denodo/lib/vdp-contrib" \
             --lib "/opt/denodo/lib/contrib" \
             --lib "/opt/denodo/lib/itp-contrib" \
             --lib "/opt/denodo/lib/maintenance-core" \
             --lib "/opt/denodo/lib/iebrowser-core" \
             --lib "/opt/denodo/lib/arn-index-client" \
             --lib "/opt/denodo/lib/extensions/jdbc-drivers/vdp-8.0" \
             --lib "/opt/denodo/lib/extensions/tfs" \
             --lib "/opt/denodo/extensions" \
             --lib "/opt/denodo/lib/util-contrib" \
             --class "/opt/denodo/extensions/dev/target/classes"
         else
           echo "Starting Denodo Virtual DataPort Server ... (press ENTER in case shell prompt was not automatically returned)"
           "$JAVA_BIN" -DDENODO_APP="Denodo VDP Server 8.0" -Xmx4096m -XX:+DisableExplicitGC -XX:+UseG1GC -XX:ReservedCodeCacheSize=256m  $DENODO_JRE11_OPTIONS $DENODO4E_DEBUG_CONF -Djava.system.class.loader=com.denodo.util.launcher.DenodoClassLoader \
             -DconfFile=$VDB_CONFIGURATION_FILE \
             -Djava.util.logging.manager=org.apache.logging.log4j.jul.LogManager \
             -Ddenodo.rmi.server.hostname.default \
             -Djavax.xml.xpath.XPathFactory:http://java.sun.com/jaxp/xpath/dom=com.sun.org.apache.xpath.internal.jaxp.XPathFactoryImpl \
             -Djavax.xml.transform.TransformerFactory=com.sun.org.apache.xalan.internal.xsltc.trax.TransformerFactoryImpl \
             -Djavax.xml.parsers.SAXParserFactory=com.sun.org.apache.xerces.internal.jaxp.SAXParserFactoryImpl \
             -Djavax.xml.parsers.DocumentBuilderFactory=com.sun.org.apache.xerces.internal.jaxp.DocumentBuilderFactoryImpl \
             -Djavax.xml.stream.XMLOutputFactory=com.ctc.wstx.stax.WstxOutputFactory \
             -D"com.microsoft.tfs.jni.native.base-directory=/opt/denodo/dll/vdp/tfs" \
             -DverboseMode=false \
             -Djdk.serialFilter="!org.mozilla.**;!org.apache.commons.**;" \
             -Ddenodo.launcher.excluded.jars="windows" \
             -Ddenodo.launcher.config="/opt/denodo/conf/vdp/startup.ini" \
             -classpath "$DENODO_LAUNCHER_CLASSPATH" \
             com.denodo.util.launcher.Launcher com.denodo.vdb.vdbinterface.server.VDBManagerImpl \
             --conf "/opt/denodo/conf" \
             --conf "/opt/denodo/resources/apache-tomcat/conf" \
             --conf "/opt/denodo/conf/vdp" \
             --lib "/opt/denodo/lib/vdp-client-core" \
             --lib "/opt/denodo/lib/vdp-server-core" \
             --lib "/opt/denodo/lib/vdp-contrib" \
             --lib "/opt/denodo/lib/contrib" \
             --lib "/opt/denodo/lib/itp-contrib" \
             --lib "/opt/denodo/lib/maintenance-core" \
             --lib "/opt/denodo/lib/iebrowser-core" \
             --lib "/opt/denodo/lib/arn-index-client" \
             --lib "/opt/denodo/lib/extensions/jdbc-drivers/vdp-8.0" \
             --lib "/opt/denodo/lib/extensions/tfs" \
             --lib "/opt/denodo/extensions" \
             --lib "/opt/denodo/lib/util-contrib" \
             --class "/opt/denodo/extensions/dev/target/classes"
         fi
        ;;
      shutdown)
        echo "Stopping Denodo VQLServer ..."
        # Load arguments list as --arg safe --arg silent
        shift
        ARGS=""
        for var in "$@"
            do
                ARGS+=" --arg $var"
            done
        "$JAVA_BIN"   $DENODO_JRE11_OPTIONS -Djava.system.class.loader=com.denodo.util.launcher.DenodoClassLoader \
          -classpath "$DENODO_LAUNCHER_CLASSPATH" \
          -DconfFile=$VDB_CONFIGURATION_FILE \
          -Djava.util.logging.manager=org.apache.logging.log4j.jul.LogManager \
          -DverboseMode=false \
          com.denodo.util.launcher.Launcher \
          com.denodo.vdb.vdbinterface.server.Shutdown \
          $ARGS \
          --conf "/opt/denodo/conf" \
          --conf "/opt/denodo/resources/apache-tomcat/conf" \
          --conf "/opt/denodo/conf/vdp" \
          --lib "/opt/denodo/lib/vdp-client-core" \
          --lib "/opt/denodo/lib/vdp-server-core" \
          --lib "/opt/denodo/lib/vdp-contrib" \
          --lib "/opt/denodo/lib/contrib" \
          --lib "/opt/denodo/lib/itp-contrib" \
          --lib "/opt/denodo/lib/maintenance-core" \
          --lib "/opt/denodo/lib/iebrowser-core" \
          --lib "/opt/denodo/lib/arn-index-client" \
          --lib "/opt/denodo/lib/extensions/jdbc-drivers/vdp-8.0" \
          --lib "/opt/denodo/lib/extensions/tfs" \
          --lib "/opt/denodo/extensions" \
          --lib "/opt/denodo/lib/util-contrib" \
          --class "/opt/denodo/extensions/dev/target/classes"
        ;;
      safe_shutdown)
        if [ "$2" == "" ]
        then
            echo "Stopping Denodo VQLServer ..."
            "$JAVA_BIN"   $DENODO_JRE11_OPTIONS -Djava.system.class.loader=com.denodo.util.launcher.DenodoClassLoader \
              -classpath "$DENODO_LAUNCHER_CLASSPATH" \
              -DconfFile=$VDB_CONFIGURATION_FILE \
              -Djava.util.logging.manager=org.apache.logging.log4j.jul.LogManager \
              -DverboseMode=false \
              com.denodo.util.launcher.Launcher \
              com.denodo.vdb.vdbinterface.server.Shutdown \
              --arg safe \
              --conf "/opt/denodo/conf" \
              --conf "/opt/denodo/resources/apache-tomcat/conf" \
              --conf "/opt/denodo/conf/vdp" \
              --lib "/opt/denodo/lib/vdp-client-core" \
              --lib "/opt/denodo/lib/vdp-server-core" \
              --lib "/opt/denodo/lib/contrib" \
              --lib "/opt/denodo/lib/itp-contrib" \
              --lib "/opt/denodo/lib/maintenance-core" \
              --lib "/opt/denodo/lib/iebrowser-core" \
              --lib "/opt/denodo/lib/arn-index-client" \
              --lib "/opt/denodo/lib/extensions/jdbc-drivers/vdp-8.0" \
              --lib "/opt/denodo/lib/extensions/tfs" \
              --lib "/opt/denodo/extensions" \
              --lib "/opt/denodo/lib/util-contrib" \
              --class "/opt/denodo/extensions/dev/target/classes"
        else
            echo "Stopping Denodo VQLServer ..."
            "$JAVA_BIN"   $DENODO_JRE11_OPTIONS -Djava.system.class.loader=com.denodo.util.launcher.DenodoClassLoader \
              -classpath "$DENODO_LAUNCHER_CLASSPATH" \
              -DconfFile=$VDB_CONFIGURATION_FILE \
              -Djava.util.logging.manager=org.apache.logging.log4j.jul.LogManager \
              -DverboseMode=false \
              com.denodo.util.launcher.Launcher \
              com.denodo.vdb.vdbinterface.server.Shutdown \
              --arg safe $2\
              --conf "/opt/denodo/conf" \
              --conf "/opt/denodo/resources/apache-tomcat/conf" \
              --conf "/opt/denodo/conf/vdp" \
              --lib "/opt/denodo/lib/vdp-client-core" \
              --lib "/opt/denodo/lib/vdp-server-core" \
              --lib "/opt/denodo/lib/vdp-contrib" \
              --lib "/opt/denodo/lib/contrib" \
              --lib "/opt/denodo/lib/itp-contrib" \
              --lib "/opt/denodo/lib/maintenance-core" \
              --lib "/opt/denodo/lib/iebrowser-core" \
              --lib "/opt/denodo/lib/arn-index-client" \
              --lib "/opt/denodo/lib/extensions/jdbc-drivers/vdp-8.0" \
              --lib "/opt/denodo/lib/extensions/tfs" \
              --lib "/opt/denodo/extensions" \
              --lib "/opt/denodo/lib/util-contrib" \
              --class "/opt/denodo/extensions/dev/target/classes"
        fi
        ;;
      *)
        echo "Usage: $0 [startup|shutdown|safe_shutdown]"
        exit 1
        ;;
    esac
    exit 0
else
    echo "Unable to execute $0: Environment variable JAVA_HOME must be set and exported"
    exit 1
fi