const LookupGoalMessage = () => {
    return {
        kind: '/v1/misc/get/goal',
        load: {}
    }
}
const LookupStatsMessage = () => {
    //dummy on server
    return {
        kind: '/v1/misc/get/stats',
        load: {}
    }
}
const ProgressMessage = () => {
    //dummy on server
    return {
        kind: '/v1/misc/get/progress',
        load: {}
    }
}
const PriorityMessage = (priority) => {
    return {
        kind: '/v1/misc/priority',
        load: {
            priority: priority
        }
    }
}
const LookupGoalMessage = (goal) => {
    return {
        kind: '/v1/misc/goal',
        load: {
            mark: goal
        }
    }
}