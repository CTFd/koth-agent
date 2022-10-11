FROM python:3.8
RUN mkdir -p /opt/app
WORKDIR /opt/app

COPY example/ /opt/app/
COPY tools/ /opt/app/
RUN pip install -r requirements.txt --no-cache-dir
RUN chmod -R 755 /opt/app

EXPOSE 8000
ENTRYPOINT ["/opt/app/serve.sh"]