package utils

import "strconv"

func ResetPasswordTemplate(firstName, LastName, secret string, id int64) string {
	userId := strconv.FormatInt(id, 10)
	return `<!DOCTYPE html PUBLIC>
	<html xmlns="http://www.w3.org/1999/xhtml">
	  <head>
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta name="x-apple-disable-message-reformatting" />
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
		<meta name="color-scheme" content="light dark" />
		<meta name="supported-color-schemes" content="light dark" />
		<title></title>
		<style type="text/css" rel="stylesheet" media="all">
		/* Base ------------------------------ */
		
		@import url("https://fonts.googleapis.com/css?family=Nunito+Sans:400,700&display=swap");
		body {
		  width: 100% !important;
		  height: 100%;
		  margin: 0;
		  -webkit-text-size-adjust: none;
		}
		
		a {
		  color: #3869D4;
		}
		
		a img {
		  border: none;
		}
		
		td {
		  word-break: break-word;
		}
		
		.preheader {
		  display: none !important;
		  visibility: hidden;
		  mso-hide: all;
		  font-size: 1px;
		  line-height: 1px;
		  max-height: 0;
		  max-width: 0;
		  opacity: 0;
		  overflow: hidden;
		}
		/* Type ------------------------------ */
		
		body,
		td,
		th {
		  font-family: "Nunito Sans", Helvetica, Arial, sans-serif;
		}
		
		h1 {
		  margin-top: 0;
		  color: #333333;
		  font-size: 22px;
		  font-weight: bold;
		  text-align: left;
		}
		
		h2 {
		  margin-top: 0;
		  color: #333333;
		  font-size: 16px;
		  font-weight: bold;
		  text-align: left;
		}
		
		h3 {
		  margin-top: 0;
		  color: #333333;
		  font-size: 14px;
		  font-weight: bold;
		  text-align: left;
		}
		
		td,
		th {
		  font-size: 16px;
		}
		
		p,
		ul,
		ol,
		blockquote {
		  margin: .4em 0 1.1875em;
		  font-size: 16px;
		  line-height: 1.625;
		}
		
		p.sub {
		  font-size: 13px;
		}
		/* Utilities ------------------------------ */
		
		.align-right {
		  text-align: right;
		}
		
		.align-left {
		  text-align: left;
		}
		
		.align-center {
		  text-align: center;
		}
		
		.u-margin-bottom-none {
		  margin-bottom: 0;
		}
		/* Buttons ------------------------------ */
		
		.button {
		  background-color: #3869D4;
		  border-top: 10px solid #3869D4;
		  border-right: 18px solid #3869D4;
		  border-bottom: 10px solid #3869D4;
		  border-left: 18px solid #3869D4;
		  display: inline-block;
		  color: #FFF;
		  text-decoration: none;
		  border-radius: 3px;
		  box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16);
		  -webkit-text-size-adjust: none;
		  box-sizing: border-box;
		}
		
		.button--green {
		  background-color: #22BC66;
		  border-top: 10px solid #22BC66;
		  border-right: 18px solid #22BC66;
		  border-bottom: 10px solid #22BC66;
		  border-left: 18px solid #22BC66;
		}
		
		.button--red {
		  background-color: #FF6136;
		  border-top: 10px solid #FF6136;
		  border-right: 18px solid #FF6136;
		  border-bottom: 10px solid #FF6136;
		  border-left: 18px solid #FF6136;
		}
		
		@media only screen and (max-width: 500px) {
		  .button {
			width: 100% !important;
			text-align: center !important;
		  }
		}
		/* Attribute list ------------------------------ */
		
		.attributes {
		  margin: 0 0 21px;
		}
		
		.attributes_content {
		  background-color: #F4F4F7;
		  padding: 16px;
		}
		
		.attributes_item {
		  padding: 0;
		}
		/* Related Items ------------------------------ */
		
		.related {
		  width: 100%;
		  margin: 0;
		  padding: 25px 0 0 0;
		  -premailer-width: 100%;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		}
		
		.related_item {
		  padding: 10px 0;
		  color: #CBCCCF;
		  font-size: 15px;
		  line-height: 18px;
		}
		
		.related_item-title {
		  display: block;
		  margin: .5em 0 0;
		}
		
		.related_item-thumb {
		  display: block;
		  padding-bottom: 10px;
		}
		
		.related_heading {
		  border-top: 1px solid #CBCCCF;
		  text-align: center;
		  padding: 25px 0 10px;
		}
		/* Discount Code ------------------------------ */
		
		.discount {
		  width: 100%;
		  margin: 0;
		  padding: 24px;
		  -premailer-width: 100%;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		  background-color: #F4F4F7;
		  border: 2px dashed #CBCCCF;
		}
		
		.discount_heading {
		  text-align: center;
		}
		
		.discount_body {
		  text-align: center;
		  font-size: 15px;
		}
		/* Social Icons ------------------------------ */
		
		.social {
		  width: auto;
		}
		
		.social td {
		  padding: 0;
		  width: auto;
		}
		
		.social_icon {
		  height: 20px;
		  margin: 0 8px 10px 8px;
		  padding: 0;
		}
		/* Data table ------------------------------ */
		
		.purchase {
		  width: 100%;
		  margin: 0;
		  padding: 35px 0;
		  -premailer-width: 100%;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		}
		
		.purchase_content {
		  width: 100%;
		  margin: 0;
		  padding: 25px 0 0 0;
		  -premailer-width: 100%;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		}
		
		.purchase_item {
		  padding: 10px 0;
		  color: #51545E;
		  font-size: 15px;
		  line-height: 18px;
		}
		
		.purchase_heading {
		  padding-bottom: 8px;
		  border-bottom: 1px solid #EAEAEC;
		}
		
		.purchase_heading p {
		  margin: 0;
		  color: #85878E;
		  font-size: 12px;
		}
		
		.purchase_footer {
		  padding-top: 15px;
		  border-top: 1px solid #EAEAEC;
		}
		
		.purchase_total {
		  margin: 0;
		  text-align: right;
		  font-weight: bold;
		  color: #333333;
		}
		
		.purchase_total--label {
		  padding: 0 15px 0 0;
		}
		
		body {
		  background-color: #F2F4F6;
		  color: #51545E;
		}
		
		p {
		  color: #51545E;
		}
		
		.email-wrapper {
		  width: 100%;
		  margin: 0;
		  padding: 0;
		  -premailer-width: 100%;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		  background-color: #F2F4F6;
		}
		
		.email-content {
		  width: 100%;
		  margin: 0;
		  padding: 0;
		  -premailer-width: 100%;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		}
		/* Masthead ----------------------- */
		
		.email-masthead {
		  padding: 25px 0;
		  text-align: center;
		}
		
		.email-masthead_logo {
		  width: 94px;
		}
		
		.email-masthead_name {
		  font-size: 16px;
		  font-weight: bold;
		  color: #A8AAAF;
		  text-decoration: none;
		  text-shadow: 0 1px 0 white;
		}
		/* Body ------------------------------ */
		
		.email-body {
		  width: 100%;
		  margin: 0;
		  padding: 0;
		  -premailer-width: 100%;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		}
		
		.email-body_inner {
		  width: 570px;
		  margin: 0 auto;
		  padding: 0;
		  -premailer-width: 570px;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		  background-color: #FFFFFF;
		}
		
		.email-footer {
		  width: 570px;
		  margin: 0 auto;
		  padding: 0;
		  -premailer-width: 570px;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		  text-align: center;
		}
		
		.email-footer p {
		  color: #A8AAAF;
		}
		
		.body-action {
		  width: 100%;
		  margin: 30px auto;
		  padding: 0;
		  -premailer-width: 100%;
		  -premailer-cellpadding: 0;
		  -premailer-cellspacing: 0;
		  text-align: center;
		}
		
		.body-sub {
		  margin-top: 25px;
		  padding-top: 25px;
		  border-top: 1px solid #EAEAEC;
		}
		
		.content-cell {
		  padding: 45px;
		}
		/*Media Queries ------------------------------ */
		
		@media only screen and (max-width: 600px) {
		  .email-body_inner,
		  .email-footer {
			width: 100% !important;
		  }
		}
		
		@media (prefers-color-scheme: dark) {
		  body,
		  .email-body,
		  .email-body_inner,
		  .email-content,
		  .email-wrapper,
		  .email-masthead,
		  .email-footer {
			background-color: #333333 !important;
			color: #FFF !important;
		  }
		  p,
		  ul,
		  ol,
		  blockquote,
		  h1,
		  h2,
		  h3,
		  span,
		  .purchase_item {
			color: #FFF !important;
		  }
		  .attributes_content,
		  .discount {
			background-color: #222 !important;
		  }
		  .email-masthead_name {
			text-shadow: none !important;
		  }
		}
		
		:root {
		  color-scheme: light dark;
		  supported-color-schemes: light dark;
		}
		</style>
		<!--[if mso]>
		<style type="text/css">
		  .f-fallback  {
			font-family: Arial, sans-serif;
		  }
		</style>
	  <![endif]-->
	  </head>
	  <body>
		<span class="preheader">Use this link to reset your password. The link is only valid for 24 hours.</span>
		<table class="email-wrapper" width="100%" cellpadding="0" cellspacing="0" role="presentation">
		  <tr>
			<td align="center">
			  <table class="email-content" width="100%" cellpadding="0" cellspacing="0" role="presentation">
				<tr>
				  <td class="email-masthead">
					<a href="https://www.themonkeys.life" class="f-fallback email-masthead_name">
					The Monkeys
				  </a>
				  </td>
				</tr>
				<!-- Email Body -->
				<tr>
				  <td class="email-body" width="570" cellpadding="0" cellspacing="0">
					<table class="email-body_inner" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation">
					  <!-- Body content -->
					  <tr>
						<td class="content-cell">
						  <div class="f-fallback">
							<h1>Hi ` + firstName + ` ` + LastName + `,</h1>
							<p>You recently requested to reset your password for your The Monkeys account. Use the button below to reset it. <strong>This password reset link is only valid for the next 5 minutes.</strong></p>
							<!-- Action -->
							<table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0" role="presentation">
							  <tr>
								<td align="center">
								
								  <table width="100%" border="0" cellspacing="0" cellpadding="0" role="presentation">
									<tr>
									  <td align="center">
										<a href="https://localhost:5001/api/v1/auth/reset-password?user=` + userId + `&evpw=` + secret + `" class="f-fallback button button--green" target="_blank">Reset your password</a>
									  </td>
									</tr>
								  </table>
								</td>
							  </tr>
							</table>
							<p>If you did not request a password reset, please ignore this email or <a href="{{https://www.themonkeys.life/contact}}">contact support</a> if you have questions.</p>
							<p>Thanks,
							  <br>The Monkeys team</p>
							<!-- Sub copy -->
							
						  </div>
						</td>
					  </tr>
					</table>
				  </td>
				</tr>
			  </table>
			</td>
		  </tr>
		</table>
	  </body>
	</html>`

}

func EmailVerificationHTML(email, secret string) string {
	return `<!DOCTYPE html>
	<html>
	<head>
	
	  <meta charset="utf-8">
	  <meta http-equiv="x-ua-compatible" content="ie=edge">
	  <title>Email Confirmation</title>
	  <meta name="viewport" content="width=device-width, initial-scale=1">
	  <style type="text/css">
	  /**
	   * Google webfonts. Recommended to include the .woff version for cross-client compatibility.
	   */
	  @media screen {
		@font-face {
		  font-family: 'Source Sans Pro';
		  font-style: normal;
		  font-weight: 400;
		  src: local('Source Sans Pro Regular'), local('SourceSansPro-Regular'), url(https://fonts.gstatic.com/s/sourcesanspro/v10/ODelI1aHBYDBqgeIAH2zlBM0YzuT7MdOe03otPbuUS0.woff) format('woff');
		}
		@font-face {
		  font-family: 'Source Sans Pro';
		  font-style: normal;
		  font-weight: 700;
		  src: local('Source Sans Pro Bold'), local('SourceSansPro-Bold'), url(https://fonts.gstatic.com/s/sourcesanspro/v10/toadOcfmlt9b38dHJxOBGFkQc6VGVFSmCnC_l7QZG60.woff) format('woff');
		}
	  }
	  /**
	   * Avoid browser level font resizing.
	   * 1. Windows Mobile
	   * 2. iOS / OSX
	   */
	  body,
	  table,
	  td,
	  a {
		-ms-text-size-adjust: 100%; /* 1 */
		-webkit-text-size-adjust: 100%; /* 2 */
	  }
	  /**
	   * Remove extra space added to tables and cells in Outlook.
	   */
	  table,
	  td {
		mso-table-rspace: 0pt;
		mso-table-lspace: 0pt;
	  }
	  /**
	   * Better fluid images in Internet Explorer.
	   */
	  img {
		-ms-interpolation-mode: bicubic;
	  }
	  /**
	   * Remove blue links for iOS devices.
	   */
	  a[x-apple-data-detectors] {
		font-family: inherit !important;
		font-size: inherit !important;
		font-weight: inherit !important;
		line-height: inherit !important;
		color: inherit !important;
		text-decoration: none !important;
	  }
	  /**
	   * Fix centering issues in Android 4.4.
	   */
	  div[style*="margin: 16px 0;"] {
		margin: 0 !important;
	  }
	  body {
		width: 100% !important;
		height: 100% !important;
		padding: 0 !important;
		margin: 0 !important;
	  }
	  /**
	   * Collapse table borders to avoid space between cells.
	   */
	  table {
		border-collapse: collapse !important;
	  }
	  a {
		color: #1a82e2;
	  }
	  img {
		height: auto;
		line-height: 100%;
		text-decoration: none;
		border: 0;
		outline: none;
	  }
	  </style>
	
	</head>
	<body style="background-color: #e9ecef;">
	
	  <!-- start preheader -->
	  <div class="preheader" style="display: none; max-width: 0; max-height: 0; overflow: hidden; font-size: 1px; line-height: 1px; color: #fff; opacity: 0;">
		A preheader is the short summary text that follows the subject line when an email is viewed in the inbox.
	  </div>
	  <!-- end preheader -->
	
	  <!-- start body -->
	  <table border="0" cellpadding="0" cellspacing="0" width="100%">
	
		<!-- start logo -->
		<tr>
		  <td align="center" bgcolor="#e9ecef">
			<!--[if (gte mso 9)|(IE)]>
			<table align="center" border="0" cellpadding="0" cellspacing="0" width="600">
			<tr>
			<td align="center" valign="top" width="600">
			<![endif]-->
			<table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
			  <tr>
				<td align="center" valign="top" style="padding: 36px 24px;">
				  <a href="https://www.themonkeys.life" target="_blank" style="display: inline-block;">
					<img src="https://www.themonkeys.life/wp-content/uploads/2019/07/blogdesire-1.png" alt="Logo" border="0" width="48" style="display: block; width: 48px; max-width: 48px; min-width: 48px;">
				  </a>
				</td>
			  </tr>
			</table>
			<!--[if (gte mso 9)|(IE)]>
			</td>
			</tr>
			</table>
			<![endif]-->
		  </td>
		</tr>
		<!-- end logo -->
	
		<!-- start hero -->
		<tr>
		  <td align="center" bgcolor="#e9ecef">
			<!--[if (gte mso 9)|(IE)]>
			<table align="center" border="0" cellpadding="0" cellspacing="0" width="600">
			<tr>
			<td align="center" valign="top" width="600">
			<![endif]-->
			<table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
			  <tr>
				<td align="left" bgcolor="#ffffff" style="padding: 36px 24px 0; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; border-top: 3px solid #d4dadf;">
				  <h1 style="margin: 0; font-size: 32px; font-weight: 700; letter-spacing: -1px; line-height: 48px;">Confirm Your Email Address</h1>
				</td>
			  </tr>
			</table>
			<!--[if (gte mso 9)|(IE)]>
			</td>
			</tr>
			</table>
			<![endif]-->
		  </td>
		</tr>
		<!-- end hero -->
	
		<!-- start copy block -->
		<tr>
		  <td align="center" bgcolor="#e9ecef">
			<!--[if (gte mso 9)|(IE)]>
			<table align="center" border="0" cellpadding="0" cellspacing="0" width="600">
			<tr>
			<td align="center" valign="top" width="600">
			<![endif]-->
			<table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
	
			  <!-- start copy -->
			  <tr>
				<td align="left" bgcolor="#ffffff" style="padding: 24px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 16px; line-height: 24px;">
				  <p style="margin: 0;">Tap the button below to confirm your email address. If you didn't create an account with <a href="https://www.themonkeys.life">The Monkeys</a>, you can safely delete this email.</p>
				</td>
			  </tr>
			  <!-- end copy -->
	
			  <!-- start button -->
			  <tr>
				<td align="left" bgcolor="#ffffff">
				  <table border="0" cellpadding="0" cellspacing="0" width="100%">
					<tr>
					  <td align="center" bgcolor="#ffffff" style="padding: 12px;">
						<table border="0" cellpadding="0" cellspacing="0">
						  <tr>
							<td align="center" bgcolor="#1a82e2" style="border-radius: 6px;">
							  <a href="https://localhost:5001/api/v1/auth/verify-email?user=` + email + `&evpw=` + secret + `" target="_blank" style="display: inline-block; padding: 16px 36px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 16px; color: #ffffff; text-decoration: none; border-radius: 6px;">Click here to verify you email</a>
							</td>
						  </tr>
						</table>
					  </td>
					</tr>
				  </table>
				</td>
			  </tr>
			  <!-- end button -->
	
			  <!-- start copy -->
			  <tr>
				<td align="left" bgcolor="#ffffff" style="padding: 24px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 16px; line-height: 24px;">
				  <p style="margin: 0;">If that doesn't work, copy and paste the following link in your browser:</p>
				  <p style="margin: 0;"><a href="https://localhost:5001/api/v1/auth/verify-email?user=` + email + `&evpw=` + secret + `" target="_blank">"https://localhost:5001/api/v1/auth/verify-email?user=` + email + `&evpw=` + secret + `"</a></p>
				</td>
			  </tr>
			  <!-- end copy -->
	
			  <!-- start copy -->
			  <tr>
				<td align="left" bgcolor="#ffffff" style="padding: 24px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 16px; line-height: 24px; border-bottom: 3px solid #d4dadf">
				  <p style="margin: 0;">Cheers,<br> Paste</p>
				</td>
			  </tr>
			  <!-- end copy -->
	
			</table>
			<!--[if (gte mso 9)|(IE)]>
			</td>
			</tr>
			</table>
			<![endif]-->
		  </td>
		</tr>
		<!-- end copy block -->
	
		<!-- start footer -->
		<tr>
		  <td align="center" bgcolor="#e9ecef" style="padding: 24px;">
			<!--[if (gte mso 9)|(IE)]>
			<table align="center" border="0" cellpadding="0" cellspacing="0" width="600">
			<tr>
			<td align="center" valign="top" width="600">
			<![endif]-->
			<table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
	
			  <!-- start permission -->
			  <tr>
				<td align="center" bgcolor="#e9ecef" style="padding: 12px 24px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 14px; line-height: 20px; color: #666;">
				  <p style="margin: 0;">You received this email because we received a request for [type_of_action] for your account. If you didn't request [type_of_action] you can safely delete this email.</p>
				</td>
			  </tr>
			  <!-- end permission -->
	
			  <!-- start unsubscribe -->
			  <tr>
				<td align="center" bgcolor="#e9ecef" style="padding: 12px 24px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 14px; line-height: 20px; color: #666;">
				  <p style="margin: 0;">To stop receiving these emails, you can <a href="https://www.themonkeys.life" target="_blank">unsubscribe</a> at any time.</p>
				  <p style="margin: 0;">Paste 1234 S. Broadway St. City, State 12345</p>
				</td>
			  </tr>
			  <!-- end unsubscribe -->
	
			</table>
			<!--[if (gte mso 9)|(IE)]>
			</td>
			</tr>
			</table>
			<![endif]-->
		  </td>
		</tr>
		<!-- end footer -->
	
	  </table>
	  <!-- end body -->
	
	</body>
	</html>`
}
