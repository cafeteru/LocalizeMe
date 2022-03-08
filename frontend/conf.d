error_page 404 /index.php;

location  index {
}

location / {
  if (!-e $request_filename){
    rewrite ^(.*)$ /index.html break;
  }
}