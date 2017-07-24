extern crate serenity;
use std::env;
use serenity::client::Client;

const PREFIX: &str = "!pig-latin";

fn main() {
    let token = env::var("TOKEN").expect("token");
    let mut client = Client::new(&token);
    client.on_message(|_, message| if message.content.starts_with(PREFIX) {
                          let text = message.content.trim_left_matches(PREFIX);
                          if let Err(why) = message.reply(&pig_latin(text)) {
                              println!("Message failed to send: {}", why);
                          }
                      });

    if let Err(why) = client.start() {
        println!("Client error: {:?}", why);
    }
}

fn pig_latin(text: &str) -> String {

    let mut chunks: Vec<String> = Vec::new();
    let mut head: Option<char> = None;
    let mut tail: String = String::new();

    for ch in text.to_lowercase().chars() {
        if ch.is_alphanumeric() {
            if tail.is_empty() && head.is_none() {
                match ch {
                    'a' | 'e' | 'i' | 'o' | 'u' => tail.push(ch),
                    _ => head = Some(ch),
                }
            } else {
                tail.push(ch);
            }
        } else {
            if !tail.is_empty() {
                match head {
                    Some(x) => {
                        tail.push(x);
                        tail.push_str("ay");
                        head = None;
                    }
                    None => tail.push_str("way"),
                }
                chunks.push(tail);
                tail = String::new();
            }

            chunks.push(ch.to_string());
        }
    }
    if !tail.is_empty() {
        match head {
            Some(x) => {
                tail.push(x);
                tail.push_str("ay");
            }
            None => tail.push_str("way"),
        }
        chunks.push(tail);
    }
    chunks.into_iter().collect()
}
