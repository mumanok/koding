###
#
#  Deployer Kite for Koding
#  Author: armagankim 
#
#  This is an example kite with two methods:
#
#    - helloWorld
#    - fooBar
#
###
kite     = require "kd-kite"
manifest = require "./manifest.json"
{spawn}  = require "child_process"
https    = require 'https' 
fs       = require 'fs-extra'
os = require 'os'
path     = require 'path'
uuid = require 'node-uuid'


class LXC
  
  constructor: (@name) ->
    @pathToContainers = "/var/lib/lxc"
    console.log "......LXC constructor",  @name

  createLxc: (callback) ->
    cmd = spawn "create-lxc", [@name]
    console.log 'create lxc'
    cmd.stdout.on 'data', (data) ->
      console.log 'stdout: ', data.toString()
    cmd.stderr.on 'data', (data) -> 
      console.log 'stderr: ', data.toString()
    cmd.on 'close', (code) ->
      console.log 'create lxc process exited with code ', code
      callback()

class Deployment
  constructor: (@zipUrl, @name, @lxc) ->
    @kitePath = path.join @lxc.pathToContainers, @lxc.name, "overlay", "opt", "kites"

  downloadAndExtractKite: (callback) ->
    {kitePath, zipUrl, name} = this
    fs.mkdirsSync(kitePath) ## TODO mkdir creates 0777, exists check
    
    filepath = path.join kitePath, "#{name}.zip"
    file     = fs.createWriteStream filepath

    console.log "downloading url", zipUrl, " to ", filepath
    
    request = https.get zipUrl, (response) ->
      response.pipe file
      response.on "end", () ->
        console.log "file download complete", response.statusCode
        if response.statusCode is 200
          console.log "chdirpath", kitePath
          process.chdir kitePath
          
          cmd = spawn "unzip", ["#{name}.zip"]
          console.log 'unziping kite'
          
          cmd.stdout.on 'data', (data) ->
            #console.log 'stdout: ', data.toString()
          cmd.stderr.on 'data', (data) -> 
            #console.log 'stderr: ', data.toString()
          
          cmd.on 'close', (code) ->
            console.log 'unzip process exited with code ', code
            fs.chmodSync path.join(kitePath, "foobar.kite"), "0777"
            fs.chown kitePath, 500000, 500000
            callback()

  runKite: () ->

    kiteInLxc = path.join "/opt", "kites", "foobar.kite"
    runner    = path.join(@lxc.pathToContainers, @lxc.name, "overlay", "opt", "runner")
    fs.writeFileSync runner, "#!/bin/bash -l \n export HOME=/root \n /usr/sbin/kd.sh #{kiteInLxc} > /tmp/kd.out & \n"
    fs.chmodSync runner, "0777"
    cmd = spawn "/usr/bin/lxc-start", ["-d", "-n", @name]
    
    cmd.stdout.on 'data', (data) ->
      console.log 'stdout kite : ', data.toString()
    cmd.stderr.on 'data', (data) -> 
      console.log 'stderr kite : ', data.toString()
    cmd.on 'close', (code) ->
      console.log 'lxc-execute process exited with code ', code


manifest.name = "Deployer_#{os.hostname()}_#{uuid.v4()}"
console.log "deployer:", manifest.name

kite.worker manifest, 

  deploy: (options, callback) ->
    {zipUrl, name} = options
    lxc = new LXC(name)
    lxc.createLxc () ->
      console.log ">>>", name
      deployment = new Deployment(zipUrl, name, lxc)
      deployment.downloadAndExtractKite deployment.runKite.bind deployment

    return callback null, "Hello, I'm #{name}! This is Deployer"


