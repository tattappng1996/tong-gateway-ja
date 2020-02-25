node {
    try {
        stage 'Build Image'
        sh('sudo docker build -t tong-gateway-ja . ')

        stage 'Run application'
        sh('sudo docker run -d -p 8080:8080 tong-gateway-ja')
    } catch (err) {
        echo 'error or some shit boi!!'
    }
}