# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
#COPY cmd .
#COPY cmd2 .
COPY . .

#COPY *.go ./

RUN go build -o cmd ./cmd/main.go && go build -o data_simulator ./data_simulator/main.go

#EXPOSE 9005
COPY start.sh start.sh
CMD ["sh", "./start.sh"]
#CMD ["chmod", "+x", "./start.sh"]

#ENTRYPOINT chmod +x ./start.sha
#RUN ["chmod", "+x", "./start.sh"]
#ENTRYPOINT ["./start.sh"]
# Теперь все, что осталось сделать, это указать Docker, какую команду выполнять, когда наш образ используется для запуска контейнера.
#Делаем это CMD командой:
#CMD /cmd && /cmd2

#Команда сборки необязательно принимает --tag флаг. Этот флаг используется для маркировки изображения строковым значением,
#которое легко прочитать и распознать людям. Если вы не передадите --tag, Docker будет использовать latest значение по умолчанию.

######## docker build --tag diplom .
######## docker run --rm -it diplom
######## docker run --rm -p 9008:9008 -it diplom
######## docker run --rm -p 8006:8006 -p 9002:9002 -it diplom

#Ваш точный вывод будет отличаться, но при условии отсутствия ошибок вы должны увидеть FINISHEDстроку в выводе сборки.
#Это означает, что Docker успешно создал наш образ и присвоил ему docker-gs-ping тег.

#Чтобы получить список изображений, выполните $ docker image ls команду (или $ docker images сокращенную запись):

#Используйте команду docker image tag(или docker tag сокращение) для создания нового тега для нашего изображения.
#Эта команда принимает два аргумента; первый аргумент — это «исходное» изображение, а второй — новый тег для создания

#Давайте удалим тег, который мы только что создали.
#Для этого мы будем использовать docker image rmкоманду или сокращение docker rmi(что означает «удалить изображение»):
#docker image rm docker-gs-ping:v1.0

#Чтобы запустить образ внутри контейнера, мы используем docker run команду.
#Для этого требуется один параметр, и это имя изображения.
#Давайте запустим наш образ и убедимся, что он работает корректно. Выполните следующую команду в своем терминале.
############### docker run demo

##Чтобы опубликовать порт для нашего контейнера, мы будем использовать --publish флаг ( -p д ля краткости) в команде запуска docker.
#Формат --publish команды такой [host_port]:[container_port].
#Итак, если бы мы хотели открыть порт 8080 внутри контейнера для порта 3000 вне контейнера, мы бы перешли 3000:8080 к --publish флагу.

######### docker run --rm -p 9006:9006 -it demo интерактивный режим

######### ######### docker run --rm -p 9006:9006 -d demo запуск контейнера в автономном режиме

##COPY commands.sh /scripts/commands.sh
##RUN ["chmod", "+x", "/scripts/commands.sh"]
##ENTRYPOINT ["/scripts/commands.sh"]