# Add missing dependencies to the go.mod file
tidy:
	go mod tidy
	
# Get app dependencies defined in the go.mod file
dep:
	go get

# Build the single binary of the app
build:
	go build

# Clean up the old binary builds
clean:
	rm ./medakcess

# Start the Medakcess daemon
start:
	sudo systemctl start medakcess

# Restart the medakcess daemon
restart:
	sudo systemctl restart medakcess

# Stop the medakcess daemon
stop:
	sudo systemctl start medakcess

# Clean up old builds, get dependecies, build ad and restart the app daemon
update: tidy dep build restart

# Configure systemd and start the medakcess daemon
setup: tidy dep build
	cat medakcess.temp > medakcess.service
	sed -i 's+User=.*+User='"$$(whoami)"'+g' ./medakcess.service
	sed -i 's+WorkingDirectory=.*+WorkingDirectory='"$$(pwd)"'+g' ./medakcess.service
	sed -i 's+ExecStart=.*+ExecStart='"$$(pwd)"'/medakcess+g' ./medakcess.service
	sudo mv ./medakcess.service /etc/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl start medakcess
