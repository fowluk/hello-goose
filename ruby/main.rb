require 'sinatra'
require 'socket'

get '/' do
  erb :index, :locals => {:hostname => Socket.gethostname()}
end
