Cách chạy docker: 
--khởi động: docker-compose up -d
--shutdown: docker-compose down
--logs: docker logs <tên container>
Cách tạo server:
--hostname: postgres
--port: 5432 (là port của postgreSQL đang chạy ở docker)
--database: chat-realtime
--user: admin
--password: admin