const axios = require('axios')
const MESG = require('mesg-js').service()

const notify = ({ endpoint, icon_emoji, text, username }, { success, error }) => axios({
	url: endpoint,
	method: 'POST',
	data: {
		icon_emoji,
		text,
		username,
	}
})
	.then(({ data }) => success(data))
	.catch(message => error({ message }))

MESG.listenTask({ notify })
