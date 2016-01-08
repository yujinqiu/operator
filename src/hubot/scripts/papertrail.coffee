path = require "path"
grpc = require "grpc"
protobuf = require "protobufjs"

protodir = path.resolve(__dirname + "/../proto")
proto = protobuf.loadProtoFile(root: protodir, file: "papertrail.proto")
papertrail = grpc.loadObject(proto.ns).papertrail
address = process.env.OPERATORD_ADDRESS
if !address
  host = process.env.OPERATORD_PORT_3000_TCP_ADDR
  port = process.env.OPERATORD_PORT_3000_TCP_PORT
  address = "#{host}:#{port}"

client = new papertrail.papertrail(address, grpc.Credentials.createInsecure())

module.exports = (robot) ->

	robot.respond /papertrail SearchTODO(sr)/, (msg) ->
		client.Search {a: 1, b: 2}, (err, response) ->
			if err
				msg.send("```\nSearch error: #{err.message}\n```")
			else
				msg.send("```\n#{response.output.PlainText}\n```")

