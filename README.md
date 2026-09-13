# Termux 

Download from: https://f-droid.org/en/
then update its repositories and install Termux 
in termux run: 
## setup 
```
curl -o https://raw.githubusercontent.com/leanderziehm/termux-android-scripts/refs/heads/main/setup_termux.sh && bash setup_termux.sh
```

# Correct Request
```
curl -X POST http://localhost:8080/api/login \
     -H "Content-Type: application/json" \
     -d '{"password": "YOUR_GENERATED_PASSWORD"}'
```