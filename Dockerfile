# menggunakan official base image golang
FROM golang:1.23.1-alpine

# set working di dalam container
WORKDIR /app

# copy seluruh kode ke dalam workdir
COPY . .

# perintah untuk ngejalanin binary diatas
RUN go build -o project-akhir-exam

# expose port 8080 di container
EXPOSE 8080

# jalankan aplikasi
CMD ["./project-akhir-exam"]