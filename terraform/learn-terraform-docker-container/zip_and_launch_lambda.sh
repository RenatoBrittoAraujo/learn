cd lambda
zip lambda_function_payload.zip
mv lambda_function_payload.zip ..
cd ..
sudo terraform apply -auto-approve