#!/bin/bash

rootPath=$GOPATH
projectName=quickModbus
projectPath=$rootPath/src/github.com/muidea/$projectName
binPath=$rootPath/bin/$projectName
imageID=""
imageNamespace=muidea.ai/develop
imageVersion=latest
imageName=$imageNamespace/$(echo $projectName | tr '[:upper:]' '[:lower:]')

cleanUp()
{
    echo "cleanUp..."
    if [ -f log.txt ]; then
        rm -f log.txt
    fi

    if [ -f $projectName ]; then
        rm -f $projectName
    fi

    if [ -f "$binPath" ]; then
        rm -f "$binPath"
    fi
}

buildBin()
{
    echo "buildBin..."
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$binPath" github.com/muidea/$projectName/cmd/$projectName
    # shellcheck disable=SC2181
    if [ $? -ne 0 ]; then
        echo "buildBin failed."
        exit 1
    else
        echo "buildBin success."
    fi
}

prepareFile()
{
    echo "prepareFile..."
    if [ ! -f "$binPath" ]; then
        buildBin
        # shellcheck disable=SC2181
        if [ $? -ne 0 ]; then
            exit 1
        fi
    fi

    cp "$binPath" ./
    # shellcheck disable=SC2181
    if [ $? -ne 0 ]; then
        echo "prepareFile failed."
        exit 1
    else
        echo "prepareFile success."
    fi
}

checkImage()
{
    echo "checkImage..."
    docker images | grep "$1" | grep "$2" > log.txt
    # shellcheck disable=SC2181
    if [ $? -eq 0 ]; then
        imageID=$(tail -1 log.txt|awk '{print $3}')
    fi
}

buildImage()
{
    echo "buildImage..."
    docker build . > log.txt
    # shellcheck disable=SC2181
    if [ $? -eq 0 ]; then
        echo "buildImage success."
    else
        echo "buildImage failed."
        exit 1
    fi

    imageID=$(tail -1 log.txt|awk '{print $3}')
}

tagImage()
{
    echo "tagImage image..."
    docker tag "$1" "$2"
    # shellcheck disable=SC2181
    if [ $? -eq 0 ]; then
        echo "tagImage success."
    else
        echo "tagImage failed."
        exit 1
    fi
}

rmiImage()
{
    echo "rmiImage..."
    docker rmi "$1":"$2"
    # shellcheck disable=SC2181
    if [ $? -eq 0 ]; then
        echo "rmiImage success."
    else
        echo "rmiImage failed."
        exit 1
    fi
}

all()
{
    echo "build $projectName docker image"

    curPath=$(pwd)

    cd "$projectPath"/docker || exit

    cleanUp

    prepareFile

    checkImage "$imageName" $imageVersion
    if [ "$imageID" ]; then
        rmiImage "$imageName" $imageVersion
    fi

    buildImage

    tagImage "$imageID" "$imageName":$imageVersion

    cleanUp

    cd "$curPath" || exit
}

build()
{
    checkImage "$imageName" $imageVersion
    if [ "$imageID" ]; then
        rmiImage "$imageName" $imageVersion
    fi

    buildImage

    tagImage "$imageID" "$imageName":$imageVersion
}

action='all'
if [ "$1" ]; then
    action=$1
fi

if [ "$action" == 'prepare' ]; then
    prepareFile
elif [ "$action" == 'clean' ]; then
    cleanUp
elif [ "$action" == 'build' ]; then
    build
else
    all
fi