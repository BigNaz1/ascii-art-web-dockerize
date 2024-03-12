# ASCII-ART-WEB-DOCKERIZE
A Dockerfile for deploying a web server capable of rendering ASCII art from text.

## AUTHORS
- Abdul Aziz
- Nezar Jaberi
- Hussain Ahmed

## USAGE
Deploy the ASCII art web server in a Docker container by following these instructions:

1. Build the Docker image:
docker image build -t <IMAGE_NAME> .

2. Initiate and run the Docker container:
docker container run -p <SERVER_PORT> --detach --name <CONTAINER_NAME> <IMAGE_NAME>

3. Activate the container's web server:
docker exec -it <CONTAINER_NAME> /bin/main.go

After setup, the web server is accessible at 'http://localhost:8080'. To generate ASCII art, submit text, select an art style, and press the 'Generate' button. The generated ASCII art will appear in a specified area, with an option to download it as a text file.

## OPTIONS
The following ASCII art styles are available for generation:
- 'standard': The default ASCII art style.
- 'shadow': A shadowed or darker ASCII art style.
- 'thinkertoy': A robotic-themed ASCII art style.
- '3d': An additional ASCII art style with 3D effects.

## ALGORITHMS USED
We utilize *SEARCHING ALGORITHMS* to identify and display ASCII art characters based on user input via the POST method.

## LICENSE
This project is distributed under the Reboot License.
