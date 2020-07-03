require 'dotenv'
require 'mina/rails'
require 'zeus'
require 'zeus/mina/make'
require 'zeus/mina/docker_service'

Dotenv.load('.env.make')

# SSH
set :user, 'root'
set :forward_agent, true # /usr/bin/ssh-add -K ~/.ssh/id_rsa

set :server, 'server_meme'
set :bot, 'bot_meme'
set :domain, 'meme'
set :registry, 'registry.gitlab.com/wearethe99/meme'
set :bot_token, -> { ENV.fetch('BOT_TOKEN') }

task :provision do
  command <<~CMD
    docker service create \
      --name #{fetch(:bot)} \
      --with-registry-auth \
      --network overlay \
      --env BOT_TOKEN=#{fetch(:bot_token)} \
      #{fetch(:registry)} \
      bot
  CMD

   command <<~CMD
      docker service create \
        --name #{fetch(:server)} \
        --with-registry-auth \
        --network overlay \
        #{fetch(:registry)} \
        server
    CMD
end

task :deploy do
  invoke :make, 'docker:push'
  invoke :'docker_service:update', fetch(:server), "--force --with-registry-auth --image #{fetch(:registry)}"
  invoke :'docker_service:update', fetch(:bot), "--force --with-registry-auth --image #{fetch(:registry)}"

  command "tput -Txterm-color setaf 3; echo 'Done <#{fetch(:domain)}>'"
end
