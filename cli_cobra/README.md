## CLI Application using Cobra


### 1 - Install Cobra-CLI
```
go install github.com/spf13/cobra-cli@latest
```

### 2 - Initialize Cobra-CLI
```
go mod init example/cobra-cli
cobra-cli init
```

### 3 - Add new command
```
cobra-cli add timezone
```

### 4 - Running the application
```
go run main.go timezone EST
```

---

### References 
- [The Cobra User Guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#using-the-cobra-library)
- [List of Time Zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
