FROM ruby:2.2-slim

RUN mkdir /app
WORKDIR /app

COPY Gemfile /app
COPY Gemfile.lock /app

RUN bundle config --global frozen 1
RUN bundle install --without development test

COPY . .

CMD ruby app.rb
