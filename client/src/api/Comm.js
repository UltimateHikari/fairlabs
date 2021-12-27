import axios from "axios";
import { UnAuthRole } from "roles";

const server_url = 'http://localhost:3000'

// TODO objects.assign(context, message.load)
const Comm = {
        get: async (context, message) => {
            console.log(context)
            await axios({
                method: 'get',
                url: server_url + message.kind,
                body: JSON.stringify(Object.assign(context, message.load))
            }).then(response => response.data)
            .catch(error => console.log(error))
        },
        post: async (context, message) => {
            await axios({
                method: 'post',
                url: server_url + message.kind,
                body: JSON.stringify(Object.assign(context, message.load))
            }).then(response => console.log(response))
            .catch(error => console.log(error))
        }
}

const FContext = {
    email: 'no email',
    username: 'no username',
    group: -1,
    course: -1,
    role: UnAuthRole
}

//Todo MessageKind, MessageLoad as interfaces in Typescript
export { Comm, FContext /*, MessageKind*/};
