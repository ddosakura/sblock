package sbjs

/*
// Example generated by sblock. CAN BE DELETED.

const NewFs = require("sblock4js")
const Raw = require("./raw")

const http = require('http')
const Koa = require('koa')
const Router = require('koa-router')

const {
	send,
} = require('sblock4koa')

const [vfs, err] = NewFs(Raw)
if (err != null) {
	throw err
}

const app = new Koa()
const router = new Router({
	prefix: '/public',
})
router.get('*', async (ctx) => {
	var path = ctx.request.url
	path = path.substr(7)
	if (path === '') {
		path = '/'
	}
	if (path[path.length - 1] === '/') {
		path += 'index.html'
	}
	await send(ctx, {
		fs: vfs,
		path,
	})
})

app.use(router.routes())
app.use(router.allowedMethods())
http.createServer(app.callback()).listen(3000)

*/
