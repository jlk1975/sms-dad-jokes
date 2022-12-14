#!/usr/bin/python
import json
import os
import requests
import emoji
import pymods.sms

# Emoji setup
rofl_emoji = emoji.emojize(":rolling_on_the_floor_laughing:")
poo_emoji = emoji.emojize(":pile_of_poo:")
joke_emojis = rofl_emoji + poo_emoji

# Set the HTTP headers
headers = {
    "Content-Type": "application/json",
    "Accept": "application/json"
}

# Set the API endpoint URL
url = "https://icanhazdadjoke.com"

# Send the request to the API endpoint
response = requests.get(url, headers=headers)
data = response.json()
joke = joke_emojis + " " + data['joke']

with open('data/joke_data.json') as f:
    data = json.load(f)

# Twilio setup
account_sid = data["account_sid"]
auth_token = data["auth_token"]
twilio_from_number = data["twilio_from_number"]
contact_list = data["contacts"]

for contact_record in contact_list:
    if contact_record['send_joke'] == "True":
        sms_sid = pymods.sms.send(account_sid, auth_token, joke, contact_record, twilio_from_number)
        print(sms_sid)