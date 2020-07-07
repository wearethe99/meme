require 'dotenv'
require 'mina/rails'
require 'zeus'
require 'zeus/mina/make'
require 'zeus/mina/docker_service'

Dotenv.load('.env.make')

# SSH
set :user, 'root'
set :forward_agent, true # /usr/bin/ssh-add -K ~/.ssh/id_rsa

set :server, 'server'
set :lukashenko_bot, 'symbal_bot'
set :domain, 'meme'
set :registry, 'registry.gitlab.com/wearethe99/meme'
set :lukashenko_bot_token, -> { ENV.fetch('LUKASHENKO_BOT_TOKEN') }

task :provision do
  command <<~CMD
    docker service create \
      --name #{fetch(:lukashenko_bot)} \
      --with-registry-auth \
      --network overlay \
      --env BOT_TOKEN=#{fetch(:lukashenko_bot_token)} \
      --env BOT_TEXT="`cat assets/lukashenko.md`" \
      --env BOT_PREFIX=lukashenko \
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
  invoke :'docker_service:update', fetch(:lukashenko_bot), "--force --with-registry-auth --image #{fetch(:registry)}"

  command "tput -Txterm-color setaf 3; echo 'Done <#{fetch(:domain)}>'"
end
