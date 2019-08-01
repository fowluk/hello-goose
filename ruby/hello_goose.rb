require 'sinatra'

set :port, ENV['PORT'] unless ENV['PORT'].nil?

get '/' do
  @instance_id = ENV['CF_INSTANCE_GUID']
  @instance_index = ENV['CF_INSTANCE_INDEX']

  erb :index
end
