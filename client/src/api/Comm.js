import axios from "axios";

const server_url = 'http://localhost:3000'

// TODO objects.assign(context, message.load)
const Comm = {
        get: async (message) => {
            await axios({
                method: 'get',
                url: server_url + message.kind,
                body: JSON.stringify(message.load)
            }).then(response => response.data)
            .catch(error => console.log(error))
        },
        post: async (message) => {
            await axios({
                method: 'post',
                url: server_url + message.kind,
                body: JSON.stringify(message.load)
            }).then(response => console.log(response))
            .catch(error => console.log(error))
        }
}

//Todo MessageKind, MessageLoad as interfaces in Typescript
export { Comm /*, MessageKind*/};
