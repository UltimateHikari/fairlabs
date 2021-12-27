const AddMessage = (name, group, amount) => {
    return {
        kind: '/v1/admin/add',
        load: {
            name: name,
            group: group,
            amount: amount
        }
    }
}

const AlgoGetMessage = () => {
    return {
        kind: '/v1/admin/get/algo',
        load: {}
    }
}

const AlgoPostMessage = (algo_id) => {
    return {
        kind: '/v1/admin/algo',
        load: {
            id: algo_id
        }
    }
}

const CondGetMessage = () => {
    return {
        kind: '/v1/admin/get/cond',
        load: {}
    }
}

const CondPostMessage = (cond_id) => {
    return {
        kind: '/v1/admin/cond',
        load: {
            cond_id
        }
    }
}

export {AddMessage, AlgoGetMessage, AlgoPostMessage, CondGetMessage, CondPostMessage}