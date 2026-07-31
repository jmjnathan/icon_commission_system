interface LabelData {
  id: number;
  clientName: string;
  phone: string;
  address: string;
  city: string;
  province: string;
  postalCode: string;
}

import logo from "../assets/logo.png";

const SENDER_NAME = "Dominic' Art - Semarang";
const SENDER_PHONE = "085622173845";

export function printShippingLabel(data: LabelData) {
  const printWindow = window.open("", "_blank", "width=800,height=550");
  if (!printWindow) return;

  const html = `
         <!DOCTYPE html>
         <html>
         <head>
           <title>Label Pengiriman #${data.id}</title>
           <style>
             @page { size: A5 landscape; margin: 0; }
             * { box-sizing: border-box; margin: 0; padding: 0; }
             body {
               font-family: 'Helvetica Neue', Arial, sans-serif;
               width: 21cm;
               height: 14.8cm;
               padding: 1cm;
               display: flex;
               flex-direction: column;
             }
             .brand {
               display: flex;
               align-items: center;
               justify-content: space-between;
               padding-bottom: 12px;
               border-bottom: 2px solid #7A1F2B;
               margin-bottom: 16px;
             }
             .brand-left {
                display: flex;
                align-items: center;
                gap: 10px;
              }

             .brand-logo-wrapper {
                width: 52px;
                height: 52px;
                border: 2px solid #7A1F2B;
                border-radius: 50%;
                overflow: hidden;   
                display: flex;
                align-items: center;
                justify-content: center;
                background: #fff;
                flex-shrink: 0;
              }

              .brand-logo-img {
                width: 100%;
                height: 100%;
                object-fit: contain;
              }
             .brand-name {
               font-size: 22px;
               font-weight: 700;
               color: #3A2E1F;
               letter-spacing: 0.5px;
             }
             .brand-tagline {
               font-size: 12px;
               color: #8A7A5C;
               text-transform: uppercase;
               letter-spacing: 1px;
             }
             .label-title {
               font-size: 16px;
               color: #8A7A5C;
               text-transform: uppercase;
               letter-spacing: 1px;
               margin-bottom: 6px;
             }
             .invocation {
               text-align: right;
               font-size: 16px;
               font-weight: 700;
               color: #7A1F2B;
               line-height: 1.4;
             }
             .invocation-line {
               display: flex;
               align-items: center;
               justify-content: flex-end;
               gap: 4px;
             }
             .cross-icon {
               width: 9px;
               height: 13px;
               flex-shrink: 0;
             }
             .recipient-name {
               font-size: 32px;
               font-weight: 700;
               color: #3A2E1F;
               margin-bottom: 8px;
             }
             .recipient-phone {
               font-size: 18px;
               color: #3A2E1F;
               margin-bottom: 10px;
             }
             .recipient-address {
               font-size: 22px;
               color: #3A2E1F;
               line-height: 1.5;
             }
             .sender {
               margin-top: auto;
               align-self: flex-end;
               text-align: right;
             }
             .sender-name {
               font-size: 18px;
               font-weight: 700;
               color: #3A2E1F;
               margin-bottom: 4px;
             }
             .sender-phone {
               font-size: 14px;
               color: #3A2E1F;
             }
             .footer {
               margin-top: auto;
               padding-top: 10px;
               border-top: 1px dashed #D9CBB0;
               font-size: 9px;
               color: #B0A588;
               text-align: center;
             }
           </style>
         </head>
         <body>
           <div class="brand">
             <div class="brand-left">
              <div class="brand-logo-wrapper">
                <img src="${logo}" class="brand-logo-img" />
              </div>

              <div>
                <div class="brand-name">Dominic's Art</div>
                <div class="brand-tagline">Sub Tutela Matris</div>
              </div>
            </div>
             <div class="invocation">
               <div class="invocation-line">
                 <svg class="cross-icon" viewBox="0 0 9 13" xmlns="http://www.w3.org/2000/svg">
                   <rect x="3.5" y="0" width="2" height="13" fill="currentColor" />
                   <rect x="0" y="3" width="9" height="2" fill="currentColor" />
                 </svg>
                 J.M.J
               </div>
               <div class="invocation-line">
                 <svg class="cross-icon" viewBox="0 0 9 13" xmlns="http://www.w3.org/2000/svg">
                   <rect x="3.5" y="0" width="2" height="13" fill="currentColor" />
                   <rect x="0" y="3" width="9" height="2" fill="currentColor" />
                 </svg>
                 S.A.G
               </div>
             </div>
           </div>
     
           <div class="label-title">Kepada</div>
           <div class="recipient-name">${data.clientName}</div>
           <div class="recipient-phone">${data.phone}</div>
           <div class="recipient-address">
             ${data.address}<br />
             ${data.city}, ${data.province} ${data.postalCode}
           </div>
   
           <div class="sender">
             <div class="label-title">Dari</div>
             <div class="sender-name">${SENDER_NAME}</div>
             <div class="sender-phone">${SENDER_PHONE}</div>
           </div>
     
           <div class="footer">Dominic's Art</div>
         </body>
         </html>
       `;

  printWindow.document.write(html);
  printWindow.document.close();
  printWindow.onload = () => {
    printWindow.print();
    printWindow.onafterprint = () => printWindow.close();
  };
}
