import 'package:flutter/material.dart';

import '../service/vehicle.dart';

class UpdateView extends StatefulWidget {
  final int vehicleId;
  final int year;
  final String model;
  final String plate;

  UpdateView({
    super.key,
    required this.vehicleId,
    required this.year,
    required this.model,
    required this.plate,
  });

  @override
  State<UpdateView> createState() => _UpdateViewState();
}

class _UpdateViewState extends State<UpdateView> {
  final VehicleService _vehicleService = VehicleService();
  late TextEditingController idController;
  late TextEditingController yearController;
  late TextEditingController modelController;
  late TextEditingController plateController;

  @override
  void initState() {
    super.initState();
    idController = TextEditingController(text: widget.vehicleId.toString());
    yearController = TextEditingController(text: widget.year.toString());
    modelController = TextEditingController(text: widget.model);
    plateController = TextEditingController(text: widget.plate);
  }

  @override
  void dispose() {
    idController.dispose();
    yearController.dispose();
    modelController.dispose();
    plateController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Atualizar usuário'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              width: 100,
              child: TextFormField(
                controller: idController,
                decoration: const InputDecoration(labelText: 'Id'),
                readOnly: true,
              ),
            ),
            SizedBox(
              width: 100,
              child: TextFormField(
                controller: yearController,
                decoration: const InputDecoration(labelText: 'Ano'),
              ),
            ),
            SizedBox(
              width: 100,
              child: TextFormField(
                controller: modelController,
                decoration: const InputDecoration(labelText: 'Modelo'),
              ),
            ),
            SizedBox(
              width: 100,
              child: TextFormField(
                controller: plateController,
                decoration: const InputDecoration(labelText: 'Placa'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(5.0),
              child: ElevatedButton(
                onPressed: () async {
                  await _vehicleService.updateData(
                    id: widget.vehicleId,
                    year: widget.year,
                    model: widget.model,
                    plate: widget.plate,
                  );
                  Navigator.pop(context);
                },
                child: const Text('Enviar'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
