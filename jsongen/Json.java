package com.qb.monetization.demo;

import org.json.JSONObject;

public class Json {
	public String username;
	public String password;

	public static Json create(JSONObject root) {
		Json object = new Json();
		object.username = root.optString("username");
		object.password = root.optString("password");

		return object;
	}
}