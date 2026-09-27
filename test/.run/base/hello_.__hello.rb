#!/usr/bin/env ruby
# A ruby test script: greets, echoes its args, and reads a file given as arg.
name = ARGV.shift || "world"
puts "RUBY: hello #{name}"
ARGV.each_with_index do |a, i|
  puts "RUBY: arg[#{i}] = #{a}"
end
if File.exist?(name)
  puts "RUBY: file first line = #{File.readlines(name).first.to_s.chomp}"
end
puts "RUBY: done"
